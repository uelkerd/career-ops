// @ts-check
/** @typedef {import('./_types.js').Provider} Provider */

/** @type {Provider} */
export default {
  id: 'theirstack',
  
  detect(entry) {
    if (entry.scan_method === 'theirstack' || entry.provider === 'theirstack') return { url: 'theirstack' };
    return null;
  },
  
  async fetch(entry, ctx) {
    const apiKey = process.env.THEIRSTACK_API_KEY;
    if (!apiKey) throw new Error('THEIRSTACK_API_KEY is not set');

    const query = entry.scan_query || entry.query;
    if (!query) throw new Error('theirstack: missing query or scan_query');

    const body = JSON.stringify({
      job_title_or: [query],
      posted_at_max_age_days: 14,
      limit: 25
    });

    const url = 'https://api.theirstack.com/v1/jobs/search';
    const json = await ctx.fetchJson(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${apiKey}`
      },
      body
    });

    const jobs = json?.data || [];
    return jobs.map(j => ({
      title: j.job_title || '',
      url: j.url || '',
      company: j.company?.name || entry.name,
      location: j.location || '',
    }));
  }
};
