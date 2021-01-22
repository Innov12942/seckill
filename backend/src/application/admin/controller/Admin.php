<?php
namespace app\admin\controller;

use think\Controller;

class Admin extends Common
{
    /**
     * 登入
     */
    public function index()
    {
        return $this->fetch('index');
    }

    public function public_edit_info()
    {
        return $this->fetch('public_edit_info');
    }

    public function change_pw()
    {
        $userid = session("user_id");
        $oldpw = input('post.oldpw');
        $newpw = input('post.newpw');
        $res = db('users')->field('id,username,password,realname,status')
            ->where(array('id' => $userid,'status' => 1))
            ->find();

        $md5_salt = config('md5_salt');
        if($res["password"] != md5($oldpw.$md5_salt)){
            return $this->success("修改失败 密码错误", 'index/index');
        }
        db("users")->where("id", $userid)->update(["password" => md5($newpw.$md5_salt)]);
        // return $this->fetch("index");
        return $this->success("修改成功", 'index/index');
    }
}
