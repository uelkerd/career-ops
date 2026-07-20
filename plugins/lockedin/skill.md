---
name: lockedin
description: Prepares the LockedIn AI environment by generating the required text files for an interview screening call.
---
# Skill: lockedin — LockedIn Protocol Prep

When the user asks to prepare the "LockedIn Protocol" or set up LockedIn AI for an interview, run this skill. 
This skill generates a highly structured package of prompts and documents that the user will upload into LockedIn AI to configure it as an ultimate copilot for their live interview.

## Inputs

1. **Company name** and **role title** (required).
   - *Normalization Rule:* For filenames, `{Company}`, `{Role}`, and `{CandidateName}` MUST be PascalCase with no spaces or special characters (e.g., "DeliveryHero", "ProductManagerTech", "DenizUlker").
   - *Slug Rule:* For the directory path (`{company-slug}-{role-slug}`), values MUST be lowercase, kebab-case, with special characters removed (e.g., "delivery-hero", "product-manager-tech").
2. **Evaluation report** in `reports/` or the **Job Description** (required to extract requirements).
3. **CV** at `cv.md` (and any tailored CV HTML output).
4. **Profile** at `config/profile.yml` and `modes/_profile.md`.
5. **Interview Research** from `interview-prep/` (if available, e.g., company intel, round structure).

## Output Format & Execution

You MUST create a new directory and write physical files to the disk. **Do not simply output the content into a single markdown file.**

1. **Create the Target Directory:** Create a directory at `interview-prep/{company-slug}-{role-slug}-lockedin/`
2. **Generate the Files:** Programmatically write the following 14 files into that new directory. Use your `write_to_file` tool to save each one.

### The Configuration Markdown File (1 file)

Generate a Markdown document named `interview-prep/{company-slug}-{role-slug}-lockedin/{company-slug}-{role-slug}-lockedin.md`. 
It must contain the following setup instructions and configuration:

**1. Interview Setup**
Provide **both** options below so the user can choose.

*Option I: Custom Prompt*
- **Prompt Name:** {Role} at {Company} (max 255 chars)
- **Prompt Content:** Write a highly specific, persona-driven prompt (max 8000 chars). Tell LockedIn AI exactly what role the user is interviewing for, the core requirements of the job, the user's specific background highlights (anchored in logic/systems thinking for this profile), and how LockedIn should act as a copilot (e.g., "Provide concise, STAR-formatted bullet points", "Flag if I am talking too long", "Focus on technical depth and executive strategy").

*Option II: Quick Setup*
Recommend the closest matching pre-built scenario from the list of LockedIn Scenarios (e.g., "Product Manager", "Software Engineering", etc.).
- **Company:** {Company}
- **Position:** {Role} (max 50 chars)
- **Interview Objective:** {e.g., "Pass the screening call by proving deep adjacent expertise and slide craftsmanship."}

**2. Advanced Configuration**
*   **Image Analysis Context:**
    Provide a short snippet (max 1000 chars) for the AI: "If I share a screenshot of a slide or an OKR document, immediately analyze it for logical flow, narrative structure, and clarity, as expected by ex-McKinsey consultants."
*   **Speech Reflection:**
    Recommend the exact settings:
    - **Industry:** {e.g., Technology / Consulting}
    - **Conversation Tone:** {Serious / Corporate}
    - **Focus Area:** {e.g., Clarity and conciseness, proving expert-level depth}
    - **Coaching Instructions:** (max 500 chars) "Ensure I answer with the 'Result' first. Cut me off if I ramble. Push me to use concrete metrics."
*   **Priority Questions (Mock Interview Setup):**
    Provide 2-3 high-priority questions to guide what the AI interviewer should ask first (e.g. to test a specific weak point or critical requirement). The total length of these questions combined MUST NOT exceed 512 characters.

### The Knowledge Injection Text Files (13 files)

Generate and save the following 13 specific `.txt` files directly into the `interview-prep/{company-slug}-{role-slug}-lockedin/` folder. The file names must be exact.

**Category: Resume**
- `Resume_{Company}_CV-{CandidateName}.txt`: Text-only, comprehensive CV tailored to the role. It MUST contain the entire career timeline, education, and credentials reframed for the role, not just recent highlights.

**Category: Job Description**
- `JobDescription_{Company}_JD-{Role}.txt`: The raw text of the job description.

**Category: Training Material**
- `TrainingMaterial_{Company}_01-CHEAT-SHEET.txt`: Quick facts, core metrics, and bottom-line reminders.
- `TrainingMaterial_{Company}_07-LESSONS-DOS-DONTS.txt`: Specific "what to say" and "what NOT to say" guidelines derived from recruiter-side heuristics. MUST explicitly include a "Red Flags to Avoid" section.
- `TrainingMaterial_{Company}_09-MINDSET.txt`: Psychological framing, overcoming imposter syndrome, and physical grounding techniques.
- `TrainingMaterial_{Company}_10-STRATEGY-HACKS.txt`: Stealth strategies (e.g., Nvidia Broadcast Eye Contact, screen sharing hacks, phone interview routing).

**Category: Knowledge Bank**
- `KnowledgeBank_{Company}_05-TERMINOLOGY-BREAKDOWN.txt`: Key company/industry-specific terms, acronyms, and frameworks the candidate must use.
- `KnowledgeBank_{Company}_06-COMPANY-DEEP-DIVE.txt`: Funding, founders, recent news, and core product proposition.
- `KnowledgeBank_{Company}_08-PEOPLE-INTEL.txt`: Interviewer persona(s), if known, and how to appeal to their background.

**Category: Supporting Material**
- `SupportingMaterial_{Company}_03-STAR-STORIES.txt`: At least 5 deep, highly tailored stories explicitly following the S-T-A-R+R (Situation, Task, Action, Result, Reflection) structure, mapped to the JD's core competencies. Draw these from `interview-prep/story-bank.md` or `cv.md` and expand on their depth.
- `SupportingMaterial_{Company}_04-QA-PREP.txt`: 5 highly likely questions and bullet-point answers. (Note: The 5 Q&A prep questions move here rather than just being advanced settings).
- `SupportingMaterial_{Company}_12-QA-PAIRS-EXTENDED.txt`: A massive database of extended interview questions and exact talking tracks (if available from previous master archives).

**Category: Sales Pitch**
- `SalesPitch_{Company}_02-TALKING-POINTS.txt`: 3-4 key themes or hooks to keep coming back to (e.g., "Systems thinking over chaos").
- `SalesPitch_{Company}_11-SALES-PITCH.txt`: A cohesive value proposition paragraph summarizing why the candidate is the perfect fit.

## Execution Rules

1. **Be Precise:** LockedIn AI needs dense, high-signal text. Avoid fluff.
2. **Adhere to the Source of Truth:** Never invent candidate experience. Base all SUPPORTING MATERIAL on `cv.md` and `article-digest.md`.
3. **Save and Inform:** Save all 14 files into the target directory and present a brief summary of the generated structure to the user.
