export const users = [
    {
        name:"visitor",
        roleId:"visitor",
        password:"visitor"
    },
    {
        name:"master",
        roleId:"master",
        password:"master"
    },
    {
        name:"admin",
        roleId:"admin",
        password:"admin"
    }
]

// 模拟服务端角色对应菜单信息--超级管理员
export const menuTreeData = [
    {
        id:1,
        name:'主页',
        path:'/index',
        icon:'operation',
        children:[
            {
                id:101,
                name:'首页',
                path:'/index/dashboard',
            }]
    },
    {
        id:2,
        name:'工单',
        path:'/task',
        icon:'operation',
        children:[
            {
                id:21,
                name:'工单详情',
                path:'/task/work',
            }]
    },
    {
        id:3,
        name:'用户',
        path:'/sys',
        icon:'setting',
        children:[
            {
                id:21,
                name:'用户管理',
                path:'/sys/user',
            },
            {
                id:22,
                name:'用户权限',
                path:'/sys/roles',
            },
        ]
    },
    {
        id:4,
        name:'应用',
        path:'/app',
        icon:'menu',
        children:[
            {
                id:11,
                name:'应用用户',
                path:'/app/user',
            },
        ]
    },
    {
        id:5,
        name:'测试',
        path:'/sit',
        icon:'odometer',
        children:[
            {
                id:12,
                name:'测试环境',
                path:'/sit/user',
            },
            {
                id:13,
                name:'测试环境2',
                path:'/sit/sit2',
            },
            {
                id:14,
                name:'测试环境3',
                path:'/sit/sit3',
            },
        ]
    },
]