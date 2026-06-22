// @ts-check
/** @typedef {import('./_types.js').Provider} Provider */

let nextAvailableTime = Date.now();

async function throttle() {
  const now = Date.now();
  if (nextAvailableTime < now) {
    nextAvailableTime = now + 14500; // 250 requests/hr = 1 req every 14.4 seconds
    return;
  }
  const delay = nextAvailableTime - now;
  nextAvailableTime += 14500;
  await new Promise(resolve => setTimeout(resolve, delay));
}

/** @type {Provider} */
export default {
  id: 'serpapi',
  
  detect(entry) {
    if (entry.scan_method === 'serpapi' || entry.provider === 'serpapi') return { url: 'serpapi' };
    // Fallback: If it's a websearch query and SERPAPI_KEY is present, SerpApi can handle it.
    if (entry.scan_method === 'websearch' && process.env.SERPAPI_KEY) return { url: 'serpapi' };
    return null;
  },
  
  async fetch(entry, ctx) {
    await throttle();
    
    const apiKey = process.env.SERPAPI_KEY;
    if (!apiKey) throw new Error('SERPAPI_KEY is not set');

    const query = entry.scan_query || entry.query;
    if (!query) throw new Error('serpapi: missing query or scan_query');

    // If query contains specific site: operators, google_jobs often fails.
    // We switch to the standard google engine to parse organic results.
    const useOrganic = query.includes('site:');
    const engine = useOrganic ? 'google' : 'google_jobs';

    const url = `https://serpapi.com/search.json?engine=${engine}&q=${encodeURIComponent(query)}&api_key=${apiKey}`;
    const json = await ctx.fetchJson(url);

    if (useOrganic) {
      const jobs = json?.organic_results || [];
      return jobs.map(j => ({
        title: j.title || '',
        url: j.link || '',
        company: entry.name || 'Organic Result',
        location: j.snippet || ''
      }));
    } else {
      const jobs = json?.jobs_results || [];
      return jobs.map(j => ({
        title: j.title || '',
        url: j.share_link || j.related_links?.[0]?.link || '',
        company: j.company_name || entry.name,
        location: j.location || '',
      }));
    }
  }
};
