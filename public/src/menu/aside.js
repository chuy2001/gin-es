// 菜单 侧边栏
export default [
  { path: '/index', title: '首页', icon: 'home' },
  {
    title: '资源管理',
    icon: 'folder-o',
    children: [
      { path: '/dashboard', title: '概览' },
      { path: '/resource', title: '资源' },
      { path: '/cmdb', title: '模型' }
    ]
  }
]
