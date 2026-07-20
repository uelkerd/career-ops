export default {
  async ingest(ctx) {
    // The core functionality of this plugin is driven by the AI agent via the skill.md file.
    // We export a no-op ingest hook to satisfy the engine validation.
    return [];
  }
};
