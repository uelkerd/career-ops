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
    const cfg = entry.theirstack || {};
    const titles = Array.isArray(cfg.job_title_or) && cfg.job_title_or.length
      ? cfg.job_title_or
      : query ? [query] : null;
    if (!titles) throw new Error('theirstack: missing query, scan_query, or theirstack.job_title_or');

    const payload = {
      job_title_or: titles,
      posted_at_max_age_days: cfg.posted_at_max_age_days ?? 14,
      limit: cfg.limit ?? 25,
    };
    if (Array.isArray(cfg.job_country_code_or) && cfg.job_country_code_or.length) {
      payload.job_country_code_or = cfg.job_country_code_or;
    }
    if (typeof cfg.remote === 'boolean') payload.remote = cfg.remote;

    const url = 'https://api.theirstack.com/v1/jobs/search';
    const json = await ctx.fetchJson(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${apiKey}`
      },
      body: JSON.stringify(payload),
    });

    const jobs = json?.data || [];
    return jobs.map(j => ({
      title: j.job_title || '',
      url: j.url || '',
      // company comes back as a plain string from this endpoint, not an object
      company: (typeof j.company === 'string' ? j.company : j.company?.name) || entry.name,
      location: j.location || '',
    }));
  }
};
