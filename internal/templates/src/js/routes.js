export default [
    { path: '/', redirect: '/processes' },
    {
        path: '/processes/:id',
        name: 'processes-preview',
        component: require('./screens/processes/preview').default,
    },

    {
        path: '/processes',
        name: 'processes',
        component: require('./screens/processes/index').default,
    },
    {
        path: '/queues',
        name: 'queues',
        component: require('./screens/queues/index').default,
    },
    {
        path: '/scheduled-jobs',
        name: 'scheduled-jobs',
        component: require('./screens/scheduled-jobs/index').default,
    },

    {
        path: '/retry-jobs',
        name: 'retry-jobs',
        component: require('./screens/retry-jobs/index').default,
    },

    {
        path: '/dead-jobs',
        name: 'dead-jobs',
        component: require('./screens/dead-jobs/index').default,
    },

    // {
    //     path: '/exceptions/:id',
    //     name: 'exception-preview',
    //     component: require('./screens/exceptions/preview').default,
    // },

    // {
    //     path: '/exceptions',
    //     name: 'exceptions',
    //     component: require('./screens/exceptions/index').default,
    // },

    // {
    //     path: '/dumps',
    //     name: 'dumps',
    //     component: require('./screens/dumps/index').default,
    // },

    // {
    //     path: '/logs/:id',
    //     name: 'log-preview',
    //     component: require('./screens/logs/preview').default,
    // },

    // {
    //     path: '/logs',
    //     name: 'logs',
    //     component: require('./screens/logs/index').default,
    // },

    // {
    //     path: '/notifications/:id',
    //     name: 'notification-preview',
    //     component: require('./screens/notifications/preview').default,
    // },

    // {
    //     path: '/notifications',
    //     name: 'notifications',
    //     component: require('./screens/notifications/index').default,
    // },

    // {
    //     path: '/jobs/:id',
    //     name: 'job-preview',
    //     component: require('./screens/jobs/preview').default,
    // },

    // {
    //     path: '/jobs',
    //     name: 'jobs',
    //     component: require('./screens/jobs/index').default,
    // },

    // {
    //     path: '/batches/:id',
    //     name: 'batch-preview',
    //     component: require('./screens/batches/preview').default,
    // },

    // {
    //     path: '/batches',
    //     name: 'batches',
    //     component: require('./screens/batches/index').default,
    // },

    // {
    //     path: '/events/:id',
    //     name: 'event-preview',
    //     component: require('./screens/events/preview').default,
    // },

    // {
    //     path: '/events',
    //     name: 'events',
    //     component: require('./screens/events/index').default,
    // },

    // {
    //     path: '/cache/:id',
    //     name: 'cache-preview',
    //     component: require('./screens/cache/preview').default,
    // },

    // {
    //     path: '/cache',
    //     name: 'cache',
    //     component: require('./screens/cache/index').default,
    // },

    // {
    //     path: '/queries/:id',
    //     name: 'query-preview',
    //     component: require('./screens/queries/preview').default,
    // },

    // {
    //     path: '/queries',
    //     name: 'queries',
    //     component: require('./screens/queries/index').default,
    // },

    // {
    //     path: '/models/:id',
    //     name: 'model-preview',
    //     component: require('./screens/models/preview').default,
    // },

    // {
    //     path: '/models',
    //     name: 'models',
    //     component: require('./screens/models/index').default,
    // },



    // {
    //     path: '/commands/:id',
    //     name: 'command-preview',
    //     component: require('./screens/commands/preview').default,
    // },

    // {
    //     path: '/commands',
    //     name: 'commands',
    //     component: require('./screens/commands/index').default,
    // },

    // {
    //     path: '/schedule/:id',
    //     name: 'schedule-preview',
    //     component: require('./screens/schedule/preview').default,
    // },

    // {
    //     path: '/schedule',
    //     name: 'schedule',
    //     component: require('./screens/schedule/index').default,
    // },

    // {
    //     path: '/redis/:id',
    //     name: 'redis-preview',
    //     component: require('./screens/redis/preview').default,
    // },

    // {
    //     path: '/redis',
    //     name: 'redis',
    //     component: require('./screens/redis/index').default,
    // },

    // {
    //     path: '/monitored-tags',
    //     name: 'monitored-tags',
    //     component: require('./screens/monitoring/index').default,
    // },

    // {
    //     path: '/gates/:id',
    //     name: 'gate-preview',
    //     component: require('./screens/gates/preview').default,
    // },

    // {
    //     path: '/gates',
    //     name: 'gates',
    //     component: require('./screens/gates/index').default,
    // },

    // {
    //     path: '/views/:id',
    //     name: 'view-preview',
    //     component: require('./screens/views/preview').default,
    // },

    // {
    //     path: '/views',
    //     name: 'views',
    //     component: require('./screens/views/index').default,
    // },
];