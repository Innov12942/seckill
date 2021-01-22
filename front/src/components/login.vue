<template>
  <div>
    <input type="text" v-model="loginForm.username" placeholder="用户名"/>
    <input type="text" v-model="loginForm.password" placeholder="密码"/>
    <button @click="login">登录</button>
  </div>
</template>
 
<script>

import axios from "axios";
export default {
  data () {
    return {
      loginForm: {
        username: '',
        password: ''
      }
    };
  },
 
  methods: {
    // ...mapMutations(['changeLogin']),
    login () {
      let _this = this;
      if (this.loginForm.username === '' || this.loginForm.password === '') {
        alert('账号或密码不能为空');
      } else {
        axios({
          method: 'post',
          url: 'api',
          data: _this.loginForm,
          crossDomain: true
        }).then(res => {
          console.log('Data: ', res.data.token);
          _this.userToken = res.data.token;
          let authtk = {Authorization: _this.userToken }
          // 将用户token保存到vuex中
        //   _this.$store.changeLogin({ Authorization: _this.userToken });
          _this.$store.commit('changeLogin', authtk);
          _this.$router.push('/home');
          alert('登陆成功');
        }).catch(error => {
          alert('账号或密码错误');
          console.log(error);
        });
      }
    }
  }
};
</script>