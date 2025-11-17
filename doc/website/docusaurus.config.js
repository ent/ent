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
      copyright: `
      Copyright © ${new Date().getFullYear()}, The Ent Authors.
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
