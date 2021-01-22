<?php
namespace app\admin\controller;

use think\Controller;

class Goods extends Common
{

    public function add_goods()
    {
        return $this->fetch('add_goods');
    }

    public function remove_goods()
    {
        return $this->fetch('remove_goods');
    }

    public function do_add()
    {
        $name = input('post.name');
        $price = intval(input('post.price'));
        $expire = input('post.expire');
        $totalnum = intval(input('post.totalnum'));
        
        $expireres = strtotime($expire);
        if($expireres == false){
            return $this->success("添加失败，请检查你的输入格式", "index/index");
        }
        if($price == 0 || $totalnum == 0){
            return $this->success("添加失败，请检查你的输入数字", "index/index");
        }

        $data = ["name" => $name, "price" => $price, "expire" => $expireres, "totalnum" => $totalnum, "remain" => $totalnum];
        $newgoodid = db("goods")->insertGetId($data);
        if($newgoodid <= 0){
            return $this->success("添加到数据库失败", "index/index");
        }
        return $this->success("添加成功", 'index/index');
    }

    public function do_remove()
    {
        $gid = intval(input('post.id'));
 
        $newgoodid = db("goods")->delete($gid);
        if($newgoodid <= 0){
            return $this->success("删除失败", "index/index");
        }
        return $this->success("删除成功", 'index/index');
    }
}
