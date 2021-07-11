module.exports={
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
    "https://buttons.github.io/buttons.js",
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
    "slackChannel": "/docs/slack",
    "newsletter": "https://www.getrevue.co/profile/ent",
    "githubRepo": "https://github.com/ent/ent"
  },
  "onBrokenLinks": "log",
  "onBrokenMarkdownLinks": "log",
  "presets": [
    [
      "@docusaurus/preset-classic",
      {
        "docs": {
          "path": "../md",
          "showLastUpdateAuthor": false,
          "showLastUpdateTime": false,
          sidebarPath: require.resolve('./sidebars.js'),
        },
        "blog": {
          "path": "blog",
          "blogSidebarCount": 'ALL',
          "blogSidebarTitle": 'All our posts',
        },
        "theme": {
          "customCss": ["../src/css/custom.css"],
        }
      }
    ]
  ],
  "plugins": [],
  "themeConfig": {
    prism: {
      additionalLanguages: ['gotemplate', 'protobuf'],
    },
    algolia: {
      apiKey: "bfc8175da1bd5078f1c02e5c8a6fe782",
      indexName: "entgo",
    },
    colorMode: {
      disableSwitch: false,
    },
    googleAnalytics: {
      trackingID: 'UA-189726777-1',
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
          href: '/docs/slack',
          position: 'right',
          className: 'header-slack-link',
          'aria-label': 'Slack channel',
        },
        {
          href: 'https://www.getrevue.co/profile/ent',
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
            {"label": "Slack", "to": "/docs/slack"},
            {"label": "Newsletter", "to": "https://www.getrevue.co/profile/ent"},
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
        src: 'https://docusaurus.io/img/oss_logo.png',
        href: 'https://opensource.facebook.com/',
      },
      copyright: `
      Copyright © ${new Date().getFullYear()} Facebook, Inc.
      The Go gopher was designed by <a href="http://reneefrench.blogspot.com/">Renee French</a>.
      <br/>
      The design is licensed under the Creative Commons 3.0 Attributions license. Read this 
      <a href="https://blog.golang.org/gopher">article</a> for more details.
      <br/>
      Design by Moriah Rich, illustration by Ariel Mashraki.
      `,
    },
    "algolia": {
      "apiKey": "bfc8175da1bd5078f1c02e5c8a6fe782",
      "indexName": "entgo"
    },
    "gtag": {
      "trackingID": "UA-189726777-1"
    },
    announcementBar: {
      id: 'user-survey', // Identify this message.
      content: 'Help us improve by taking <a href="https://forms.gle/vLUXn2ETDD2q457X9">our user survey.</a>',
      backgroundColor: '#fafbfc',
      textColor: '#091E42',
      isCloseable: true,
    },
  }
}