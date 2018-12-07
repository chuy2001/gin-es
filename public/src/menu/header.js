// 菜单 顶栏
export default [
  { path: '/index', title: '首页', icon: 'home' },
  {
    title: '资源管理',
    icon: 'folder-o',
    children: [
      { path: '/resource', title: '资源' },
      { path: '/cmdb', title: '模型' }
    ]
  }
]
