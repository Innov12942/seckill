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
          data: `{"token":"` + this.$store.state.Authorization + `","getresult":"true"}`,
          crossDomain: true
        }).then(res => {
          console.log('Data: ', res.data);
          var goodsstr = res.data.toString();
          goodsstr = goodsstr.substring(0, goodsstr.length - 1);

             console.log('array: ', goodsstr);
            if (goodsstr.indexOf("&") != -1){
            var goodarray = goodsstr.split('&');

                goodarray.forEach((element, index) => {
                    var ul = document.getElementById("ulcontent");
                    var li = document.createElement("li");
                    li.setAttribute("id", index + 100);
                    var obj = JSON.parse(element);


                    li.innerHTML = "已经秒杀成功商品 : [ " + obj.Gid + " ]";
                    ul.appendChild(li);
                });

                alert('获得秒杀商品数据成功');
            }else if(goodsstr.indexOf("i") != -1){
                var obj = JSON.parse(goodsstr);
                var ul = document.getElementById("ulcontent");
                var li = document.createElement("li");
                li.innerHTML = "已经秒杀成功商品 : [ " + obj.Gid + " ]";
                ul.appendChild(li);
                alert('获得秒杀商品数据成功');
            }else{
                alert("无秒杀成功商品");
            };

          // 将用户token保存到vuex中
        //   _this.$store.changeLogin({ Authorization: _this.userToken });
        }).catch(error => {
          alert('获得秒杀商品数据失败');
          console.log(error);
          this.$router.push('/login');
        });
      }
    
  
};
</script>