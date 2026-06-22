// @ts-check
/** @typedef {import('./_types.js').Provider} Provider */

/** @type {Provider} */
export default {
  id: 'jsearch',
  
  detect(entry) {
    if (entry.scan_method === 'jsearch' || entry.provider === 'jsearch') return { url: 'jsearch' };
    return null;
  },
  
  async fetch(entry, ctx) {
    const apiKey = process.env.JSEARCH_API_KEY || process.env.RAPIDAPI_KEY;
    if (!apiKey) throw new Error('JSEARCH_API_KEY or RAPIDAPI_KEY is not set');
    
    const query = entry.scan_query || entry.query;
    if (!query) throw new Error('jsearch: missing query or scan_query');

    const url = `https://jsearch.p.rapidapi.com/search?query=${encodeURIComponent(query)}&page=1&num_pages=1`;
    const json = await ctx.fetchJson(url, {
      headers: {
        'x-rapidapi-host': 'jsearch.p.rapidapi.com',
        'x-rapidapi-key': apiKey
      }
    });

    const jobs = json?.data || [];
    return jobs.map(j => ({
      title: j.job_title || '',
      url: j.job_apply_link || j.job_google_link || '',
      company: j.employer_name || entry.name,
      location: j.job_city ? `${j.job_city}, ${j.job_country}` : (j.job_country || ''),
    }));
  }
};
