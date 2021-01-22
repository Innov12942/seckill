<template>
  <div>
    <ul id="ulcontent">

    </ul>
  </div>
</template>
 
<script>

import axios from "axios";
export default {
  data () {
    return {
    };
  },
 

    // ...mapMutations(['changeLogin']),
    created () {
        axios({
          method: 'post',
          url: 'api',
          data: `{"token":"` + this.$store.state.Authorization + `","getgoods":"true"}`,
          crossDomain: true
        }).then(res => {
          console.log('Data: ', res.data);
          var goodsstr = res.data.toString();
          goodsstr = goodsstr.substring(0, goodsstr.length - 1);

          var goodarray = goodsstr.split('&');
          console.log('Array: ', goodarray);
          // console.log('type: ' + typeof goodsstr);


          goodarray.forEach((element, index) => {
            var ul = document.getElementById("ulcontent");
            var li = document.createElement("li");
            li.setAttribute("id", index);
            var obj = JSON.parse(element);

            var unix_timestamp = obj.Expire
              var a = new Date(unix_timestamp * 1000);
              var year = a.getFullYear();
              var month = a.getMonth() + 1;
              var date = a.getDate();
              var hour = a.getHours();
              var min = a.getMinutes();
              var sec = a.getSeconds();
              var time = year + '-' + month + '-'+ date + ' ' + hour + ':' + min + ':' + sec ;

            li.innerHTML = "商品名称 : [" + obj.Name + "] 商品价格 : [" + obj.Price + "] 商品总数量 : [" + obj.Totalnum +  "] 商品秒杀时间 : [" + time + "]";
            ul.appendChild(li);

            var newBtn = document.createElement('button');
            newBtn.innerHTML = "秒杀"
            newBtn.id = index
            newBtn.onclick = () => { 
              var d = new Date();
              var curtime = d.getTime();
              if(curtime < unix_timestamp * 1000){
                alert("秒杀时间未到");
                return;
              }
              axios({
                method : 'post',
                url : "api",
                data: `{"token":"` + this.$store.state.Authorization + `","goodid":"` + obj.Id + `"}`
              }).then(
                res => {
                  console.log("click button");
                  alert('秒杀请求已提交，请勿重复提交');
                  this.$router.push('/result');
                  
                }
              ).catch(error => {
                console.log(error);
              });

            }
            li.appendChild(newBtn);
          });

          // 将用户token保存到vuex中
        //   _this.$store.changeLogin({ Authorization: _this.userToken });
          alert('获得商品数据成功');
        }).catch(error => {
          alert('获得商品数据失败');
          console.log(error);
          this.$router.push('/login');
        });
      }
    
  
};
</script>