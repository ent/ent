module.exports = {
  md: [
    {
      type: 'category',
      label: 'Getting Started',
      items: [
          'getting-started',
      ],
      collapsed: false,
    },
    {
      type: 'category',
      label: 'Schema',
      items: [
        'schema-def',
        'schema-fields',
        'schema-edges',
        'schema-indexes',
        'schema-mixin',
        'schema-annotations',
      ],
      collapsed: false,
    },
    {
      type: 'category',
      label: 'Code Generation',
      items: [
        'code-gen',
        'crud',
        'traversals',
        'eager-load',
        'hooks',
        'privacy',
        'transactions',
        'predicates',
        'aggregate',
        'paging',
      ],
      collapsed: false,
    },
    {
      type: 'category',
      label: 'Migration',
      items: [
        'migrate',
        'dialects',
      ],
      collapsed: false,
    },
    {
      type: 'category',
      label: 'Misc',
      items: [
        'templates',
        'graphql',
        'sql-integration',
        'testing',
        'faq',
        'feature-flags',
        'translations',
        'contributors'
      ],
      collapsed: false,
    },
  ],
  tutorial: [
    {
      type: 'category',
      label: 'First Steps',
      items: [
        'tutorial-setup',
        'tutorial-todo-crud',
      ],
      collapsed: false,
    },
    {
      type: 'category',
      label: 'GraphQL Basics',
      items: [
        'tutorial-todo-gql',
        'tutorial-todo-gql-node',
        'tutorial-todo-gql-paginate',
        'tutorial-todo-gql-field-collection',
        'tutorial-todo-gql-tx-mutation',
        'tutorial-todo-gql-mutation-input',
      ],
      collapsed: false,
    },
  ]
}
