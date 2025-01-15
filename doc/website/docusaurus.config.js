const TwitterSvg =
    '<svg style="fill: #1DA1F2; vertical-align: middle;" width="16" height="16" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M459.37 151.716c.325 4.548.325 9.097.325 13.645 0 138.72-105.583 298.558-298.558 298.558-59.452 0-114.68-17.219-161.137-47.106 8.447.974 16.568 1.299 25.34 1.299 49.055 0 94.213-16.568 130.274-44.832-46.132-.975-84.792-31.188-98.112-72.772 6.498.974 12.995 1.624 19.818 1.624 9.421 0 18.843-1.3 27.614-3.573-48.081-9.747-84.143-51.98-84.143-102.985v-1.299c13.969 7.797 30.214 12.67 47.431 13.319-28.264-18.843-46.781-51.005-46.781-87.391 0-19.492 5.197-37.36 14.294-52.954 51.655 63.675 129.3 105.258 216.365 109.807-1.624-7.797-2.599-15.918-2.599-24.04 0-57.828 46.782-104.934 104.934-104.934 30.213 0 57.502 12.67 76.67 33.137 23.715-4.548 46.456-13.32 66.599-25.34-7.798 24.366-24.366 44.833-46.132 57.827 21.117-2.273 41.584-8.122 60.426-16.243-14.292 20.791-32.161 39.308-52.628 54.253z"></path></svg>';

const config = {
  "title": "ent",
  "i18n": {
    "defaultLocale": 'en',
    "locales": ['en', 'zh', 'ja', 'he'],
    "localeConfigs": {
      "en": {
        "label": 'English',
        "direction": 'ltr',
      },
      "zh": {
        "label": 'Chinese',
        "direction": 'ltr',
      },
      "ja": {
        "label": 'Japanese',
        "direction": 'ltr',
      },
      "he": {
        "label": 'Hebrew',
        "direction": 'rtl',
      },
    },
  },
  "tagline": "An entity framework for Go",
  "url": "https://entgo.io",
  "baseUrl": "/",
  "organizationName": "ent",
  "projectName": "ent",
  "scripts": [
    "https://cdnjs.cloudflare.com/ajax/libs/clipboard.js/2.0.0/clipboard.min.js",
    "/js/code-block-buttons.js",
    "/js/custom.js"
  ],
  "favicon": "img/favicon.ico",
  "customFields": {
    "users": [
      {
        "caption": "User1",
        "image": "/img/undraw_open_source.svg",
        "infoLink": "https://www.facebook.com",
        "pinned": true
      }
    ],
    "slackChannel": "/docs/community#slack",
    "newsletter": "https://entgo.substack.com/",
    "githubRepo": "https://github.com/ent/ent"
  },
  "onBrokenLinks": "log",
  "onBrokenMarkdownLinks": "log",
  "presets": [
    [
      "@docusaurus/preset-classic",
      {
        "docs": {
          path: "../md",
          editUrl: 'https://github.com/ent/ent/edit/master/doc/md/',
          showLastUpdateAuthor: true,
          showLastUpdateTime: true,
          sidebarPath: require.resolve('./sidebars.js'),
        },
        gtag: {
          trackingID: "UA-189726777-1",
        },
        "blog": {
          "feedOptions": {
            "type": 'all',
            "copyright": `Copyright © ${new Date().getFullYear()}, The Ent Authors.`,
          },
          "path": "blog",
          "blogSidebarCount": 'ALL',
          "blogSidebarTitle": 'All our posts',
        },
        "theme": {
          "customCss": require.resolve('./src/css/custom.css'),
        }
      }
    ]
  ],
  "plugins": [
      [
        `@docusaurus/plugin-client-redirects`,{
          redirects: [
            {
              to:'/docs/community',
              from:'/docs/slack',
            },
          ],
        },
      ]
  ],
  "themeConfig": {
    prism: {
      additionalLanguages: ['gotemplate', 'protobuf', "hcl"],
      magicComments: [
        {
          className: 'theme-code-block-highlighted-line',
          line: 'highlight-next-line',
          block: {start: 'highlight-start', end: 'highlight-end'},
        },
        {
          className: 'code-block-error-message',
          line: 'highlight-next-line-error-message',
        },
        {
          className: 'code-block-info-line',
          line: 'highlight-next-line-info',
          block: {start: 'highlight-info-start', end: 'highlight-info-end'},
        },
      ],
    },
    algolia: {
      appId: "8OIT9XHKR1",
      apiKey: "42c78b88ab39bda9adad782eba9e2aa2",
      indexName: "entgo",
    },
    colorMode: {
      disableSwitch: false,
    },
    "navbar": {
      "title": "",
      "logo": {
        "src": "img/logo.png"
      },
      "items": [
        {
          "to": "docs/getting-started",
          "label": "Docs",
          "position": "left"
        },
        {
          "to": "docs/tutorial-setup",
          "label": "Tutorials",
          "position": "left"
        },
        {
          "href": "https://pkg.go.dev/entgo.io/ent?tab=doc",
          "label": "GoDoc",
          "position": "left",
          "className": "header-godoc-link",
        },
        {to: 'blog', label: 'Blog', position: 'left'},
        {
          href: '/docs/community#slack',
          position: 'right',
          className: 'header-slack-link',
          'aria-label': 'Slack channel',
        },
        {
          href: 'https://discord.gg/qZmPgTE6RX',
          position: 'right',
          className: 'header-discord-link',
          'aria-label': 'Discord Server',
        },
        {
          href: 'https://entgo.substack.com/',
          position: 'right',
          className: 'header-newsletter-link',
          'aria-label': 'Newsletter page',
        },
        {
          href: 'https://twitter.com/entgo_io',
          position: 'right',
          className: 'header-twitter-link',
          'aria-label': 'Twitter page',
        },
        {
          href: 'https://github.com/ent/ent',
          position: 'right',
          className: 'header-github-link',
          'aria-label': 'GitHub repository',
        },
        {
          type: 'localeDropdown',
          position: 'right',
          dropdownItemsAfter: [
            {
              to: '/docs/translations',
              label: 'Help Us Translate',
            },
          ],
        },
      ]
    },
    "image": "img/undraw_online.svg",
    ogImage: 'img/undraw_online.svg',
    twitterImage: 'img/undraw_tweetstorm.svg',
    "footer": {
      "links": [
        {
          "title": "Docs",
          "items": [
            {"label": "Getting Started", "to": "/docs/getting-started"},
            {"label": "Schema Guide", "to": "/docs/schema-def"},
            {"label": "Code Generation", "to": "/docs/code-gen"},
            {"label": "GraphQL Integration", "to": "/docs/graphql"},
            {"label": "Schema Migration", "to": "/docs/migrate"},
          ]
        },
        {
          "title": "Community",
          "items": [
            {"label": "GitHub", "to": "https://github.com/ent/ent"},
            {"label": "Slack", "to": "/docs/community#slack"},
            {"label": "Discord", "to": "https://discord.gg/qZmPgTE6RX"},
            {"label": "Newsletter", "to": "https://entgo.substack.com/"},
            {"label": "Discussions", "to": "https://github.com/ent/ent/discussions"},
            {
              "label": "Twitter",
              "to": "https://twitter.com/entgo_io"
            }
          ]
        },
        {
          "title": "Legal",
          "items": [
            {"label": "Privacy", "to": "https://opensource.facebook.com/legal/privacy/"},
            {"label": "Terms", "to": "https://opensource.facebook.com/legal/terms/"},
            {"label": "Data Policy", "to": "https://opensource.facebook.com/legal/data-policy/"},
            {"label": "Cookie Policy", "to": "https://opensource.facebook.com/legal/cookie-policy/"},
          ]
        },
        {
          "title": "Social",
          "items": [
            {"html": `
            <a href="https://github.com/ent/ent/stargazers">
                <img src="https://img.shields.io/github/stars/ent/ent?style=social"/>
            </a>`},
            {"html": `
            <a href="https://twitter.com/entgo_io">
                <img src="https://img.shields.io/twitter/follow/entgo_io?style=social"/>
            </a>`}
          ]
        }
      ],
      logo: {
        alt: 'Facebook Open Source Logo',
        src: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAVQAAABaCAMAAADQMV5DAAAC3FBMVEUAAAD///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////+GTFYVAAAA9HRSTlMAG7O/KQR48v/+hA5W2+piObG8RBSU/fehIG7stSWl9nkIStbYU0XN4FUvrHoDcOi3OAuLmheOmRVm4rtBNbDucsnfXlDUJpuGB3vtM3+mHBKMEF3FSDzmaMLpWN7PRh2XkAaAo3X4riuiBctS3WAxtvFx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',
        href: 'https://opensource.facebook.com/',
      },
      copyright: `
      Copyright © ${new Date().getFullYear()} Facebook, Inc.
      The Go gopher was designed by <a href="http://reneefrench.blogspot.com/">Renee French</a>.
      <br/>
      The design is licensed under the Creative Commons 3.0 Attributions license. Read this 
      <a href="https://go.dev/blog/gopher">article</a> for more details.
      <br/>
      Design by Moriah Rich, illustration by Ariel Mashraki.
      `,
    },
    announcementBar: {
      id: 'announcementBar-2', // Increment on change
      // content: `⭐️ If you like Ent, give it a star on <a target="_blank" rel="noopener noreferrer" href="https://github.com/ent/ent">GitHub</a> and follow us on <a target="_blank" rel="noopener noreferrer" href="https://twitter.com/entgo_io" >Twitter</a> ${TwitterSvg}`,
      content: `<a style="text-decoration: none;" target="_blank" rel="noopener noreferrer" href="https://twitter.com/Israel/status/1713931519620788538">The Ent Team Stands With Israel 🇮🇱</a>`,
      backgroundColor: '#fafbfc',
      textColor: '#404756',
      isCloseable: false,
    },
  }
};

module.exports = config;
