// @ts-check
/** @typedef {import('./_types.js').Provider} Provider */

/** @type {Provider} */
export default {
  id: 'universal_feed',
  
  detect(entry) {
    if (entry.scan_method === 'universal_feed' || entry.provider === 'universal_feed') return { url: 'universal_feed' };
    return null;
  },
  
  async fetch(entry, ctx) {
    const url = entry.feed_url || entry.api || entry.careers_url;
    if (!url) throw new Error('universal_feed: missing feed_url or careers_url');

    let responseText = await ctx.fetchText(url);
    
    // Attempt to parse as JSON first
    let isJson = false;
    let jsonObj = null;
    try {
      jsonObj = JSON.parse(responseText);
      isJson = true;
    } catch(e) {
      isJson = false;
    }

    const jobs = [];

    if (isJson) {
      // Try to find the array of jobs dynamically
      let arr = [];
      if (Array.isArray(jsonObj)) {
        arr = jsonObj;
      } else if (jsonObj && typeof jsonObj === 'object') {
        // Look for common array fields or just grab the first array we find
        for (const key of ['data', 'jobs', 'results', 'items', 'offers']) {
          if (Array.isArray(jsonObj[key])) {
            arr = jsonObj[key];
            break;
          }
        }
        if (!arr.length) {
          for (const key of Object.keys(jsonObj)) {
            if (Array.isArray(jsonObj[key])) {
              arr = jsonObj[key];
              break;
            }
          }
        }
      }

      if (!arr.length) throw new Error('universal_feed: Could not locate job array in JSON response');

      const titleField = entry.title_field || 'title';
      const urlField = entry.url_field || 'url';
      const companyField = entry.company_field || 'company';

      for (const item of arr) {
        // Fallbacks for common field names if the user didn't specify mapping
        const title = item[titleField] || item.job_title || item.name || item.position || '';
        const link = item[urlField] || item.link || item.apply_url || item.absolute_url || '';
        const company = item[companyField] || item.company_name || entry.name || '';
        
        if (title && link) {
          jobs.push({ title, url: link, company, location: item.location || '' });
        }
      }
    } else {
      // Parse as XML (RSS/Atom) using regex (since no DOMParser is available in basic Node.js)
      const itemRegex = /<(?:item|entry)[\s\S]*?>([\s\S]*?)<\/(?:item|entry)>/gi;
      let match;
      while ((match = itemRegex.exec(responseText)) !== null) {
        const itemXml = match[1];
        
        const titleMatch = itemXml.match(/<title[^>]*>([\s\S]*?)<\/title>/i);
        let title = titleMatch ? titleMatch[1].replace(/<!\[CDATA\[(.*?)\]\]>/gi, '$1').trim() : '';

        let link = '';
        const linkHrefMatch = itemXml.match(/<link[^>]*href=["']([^"']+)["']/i);
        const linkTagMatch = itemXml.match(/<link[^>]*>([\s\S]*?)<\/link>/i);
        
        if (linkHrefMatch) {
          link = linkHrefMatch[1];
        } else if (linkTagMatch) {
          link = linkTagMatch[1].replace(/<!\[CDATA\[(.*?)\]\]>/gi, '$1').trim();
        }

        const companyMatch = itemXml.match(/<(?:company|author|dc:creator)[^>]*>([\s\S]*?)<\/(?:company|author|dc:creator)>/i);
        let company = companyMatch ? companyMatch[1].replace(/<!\[CDATA\[(.*?)\]\]>/gi, '$1').replace(/<[^>]*>/g, '').trim() : entry.name;

        if (title && link) {
          jobs.push({ title, url: link, company, location: '' });
        }
      }
    }

    return jobs;
  }
};
