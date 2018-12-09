import Mock from 'mockjs'

const userDB = [{
  username: 'admin',
  password: 'admin',
  uuid: 'admin-uuid',
  name: '管理员'
},
{
  username: 'editor',
  password: 'editor',
  uuid: 'editor-uuid',
  name: '编辑'
},
{
  username: 'user1',
  password: 'user1',
  uuid: 'user1-uuid',
  name: '用户1'
}
]

Mock.mock('/api/login', 'post', ({ url, type, body }) => {
  const bodyObj = JSON.parse(body)
  const user = userDB.find(e => e.username === bodyObj.username && e.password === bodyObj.password)
  if (user) {
    return {
      code: 0,
      msg: '登录成功',
      data: {
        ...user,
        token: 'd787syv8dys8cas80d9s0a0d8f79ads56f7s4d56f879a8as89fd980s7dg'
      }
    }
  } else {
    return {
      code: 401,
      msg: '用户名或密码错误',
      data: {}
    }
  }
})

Mock.mock('/v1/mgmt/table', 'get', ({ url, type, body }) => {
  return Mock.mock([{
    'id|+1': 1,
    'name': '@CNAME',
    'alias': '@NAME',
    'creator_username': '@CNAME',
    'creation_time': '@DATE'
  }, {
    'id|+1': 1,
    'name': '@CNAME',
    'alias': '@NAME',
    'creator_username': '@CNAME',
    'creation_time': '@DATE',
    'fields': [{
      'name': 'id',
      'alias': 'id',
      'readme': 'id',
      'type': 0,
      'is_multi': false,
      'required': false
    }]
  }])
})

Mock.mock('/v1/mgmt/table', 'put', ({ url, type, body }) => {
  const bodyObj = JSON.parse(body)
  console.log(bodyObj)
  return bodyObj
  // return Mock.mock({
  //   'id|+1': 1,
  //   'name': '@CNAME',
  //   'alias': '@NAME',
  //   'creator_username': '@CNAME',
  //   'creation_time': '@DATE'
  // })
})

Mock.mock('/v1/mgmt/table', 'delete', () => {
  return {
    code: 0
  }
})

Mock.mock(RegExp('/v2/mgmt/instance' + '.*'), 'post', ({ body }) => {
  // 这是通过 post 传来的参数
  body = JSON.parse(body)
  console.log('模拟请求', body)
  const { page } = body
  page.total = 1000
  return Mock.mock(
    {
      code: 0,
      msg: '获取数据成功',
      data: {
        page,
        'list|20': [
          {
            'id': '@guid',
            'value|1': [10, 100, 200, 500],
            'type': '@boolean',
            'admin': '@cname',
            'adminNote': '@cparagraph(0.5)',
            'dateTimeCreat': '@datetime',
            'used': '@boolean',
            'dateTimeUse': '@datetime'
          }
        ]
      }
    }
  )
})
