<?php
namespace app\admin\controller;

use think\Controller;

class Index extends Common
{
    /**
     * 登入
     */
    public function index()
    {
        $where = array();
        $goodsinfo = db('goods')->where($where)->select();
        $this->assign('goodsinfo', $goodsinfo);
        return $this->fetch('index');
    }
}
