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

// Mock.mock('/api/v1/mgmt/table', 'post', ({ url, type, body }) => {
//   const bodyObj = JSON.parse(body)
//   console.log(bodyObj)
//   bodyObj['creator_username'] = 'admin'
//   bodyObj['creation_time'] = formatDate (new Date())
//   return bodyObj
// })

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

var  formatDate  =   function (date)  {      
  var y  =  date.getFullYear()     
  var m  =  date.getMonth() +1; 
  m = m <10 ? '0'+m:m    
  var  d  =  date.getDate(); 
  d  = d <10 ? ('0'+ d): d    
  return y + '-' + m + '-' + d
}
