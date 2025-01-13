const businessRoutes = [
    {
        path: '/index/dashboard', // 作为 /home/dashboard
        name: 'dashboard',
        component: () => import('@/views/index/dashboard.vue'),
    },
    {
        path: '/task/work', // 直接以 /task/work 访问
        name: 'task',
        component: () => import('@/views/task/work.vue'),
    },
    {
        path: '/sys/user',
        name: 'sys_user',
        component: () => import('@/views/sys/user.vue'),
    },
    {
        path: '/sys/roles',
        name: 'sys_roles',
        component: () => import('@/views/sys/roles.vue'),
    },
    {
        path: '/app/user',
        name: 'app_user',
        component: () => import('@/views/app/user.vue'),
    },
    {
        path: '/sit/user',
        name: 'sit_user',
        component: () => import('@/views/sit/user.vue'),
    },
    {
        path: '/sys',
        redirect: '/sys/user', // 默认重定向到 sys_user
    },
    {
        path: '/app',
        redirect: '/app/user', // 默认重定向到 app_user
    },
];

export default businessRoutes;  