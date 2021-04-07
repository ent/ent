---
id: translations
title: Translations
---

## Introduction

To make Ent more accessible to native speakers of other languages we have started a translation initiative.
Our goal is to make all documents in this website available in both Chinese and Japanese. To facilitate the 
translation, approval and deployment process we have integrated this website with [Crowdin](https://crowdin.com).

## Contributing

1. Sign up as a translator on our [Crowdin project](https://crwd.in/ent)
2. After registration and sign up, go to our [Project Page](https://crowdin.com/project/ent)
3. Pick the language you would like to translate a document into.
4. Pick the document you want to translate:
   * All Docs are under the `md/` folder
   * Blog Posts are under `website/blog`
   * Website UI components are under `website/i18n`
5. For more information on using the Crowdin UI, please read the [Online Editor documentation](https://support.crowdin.com/online-editor/).
6. Your suggestions will be reviewed by a proof-reader, and once approved they will be 
included in the next website deployment. 

## Joining as a Proof-reader
If you would like to contribute in a more permanent way and become a proof-reader
for one of the languages, please ping us via the slack channel.

## Important Guidelines

- Most documents begin with a line: `id: <document-id>` - this should NEVER be translated and will actually break the website build if it is.
- The website build is sensitive to closing of HTML elements, if HTML is needed within the translation it must be closed,
  or self-closing (i.e. do not leave a dangling `<br>`, make it `<br/>`).
