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
        'versioned-migrations',
        'dialects',
      ],
      collapsed: false,
    },
    {
      type: 'category',
      label: 'Misc',
      items: [
        'templates',
        'extensions',
        'graphql',
        'sql-integration',
        'ci',
        'testing',
        'faq',
        'feature-flags',
        'generating-ent-schemas',
        'translations',
        'contributors',
        'writing-docs',
        'slack'
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
        'tutorial-todo-gql-filter-input',
        'tutorial-todo-gql-schema-generator',
      ],
      collapsed: false,
    },
    {
      type: 'category',
      collapsed: false,
      label: 'gRPC',
      items: [
          'grpc-intro',
          'grpc-setting-up',
          'grpc-generating-proto',
          'grpc-generating-a-service',
          'grpc-server-and-client',
          'grpc-edges',
          'grpc-optional-fields',
          'grpc-service-generation-options',
      ]
    }
  ]
}
