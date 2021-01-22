<?php


namespace app\admin\model;

use think\Db;

class Admin extends \think\Model {

    public $status = array(1 => '无效', 2 => '有效');

    /**
     * 登录时调用
     * @param String $username 用户名
     * @return Array
     */
    public function getInfoByUsername($username) {
        $res = db('users')->field('id,username,password,realname,status')
            ->where(array('username' => $username,'status' => 1))
            ->find();
        if ($res) {
            $res = $res;
        }

        return $res;
    }

    /**
     *
     * @param int $userid 用户ID
     * @return Array
     */
    public function getUserGroups($uid) {

        $res = db('users')->field('group')->where('id', $uid)->select();

        $userGroups = '';
        if ($res) {
            foreach ($res as $k => $v) {
                $userGroups .= $v['group'] . ',';
            }
            return trim($userGroups, ',');
        } else {
            return false;
        }
    }

    /**
     * 登陆更新
     * @param int $id id
     * @param array $data 更新的数据
     */
    public function editInfo( $id, $data = array()) {
        $data['lastlogintime'] = time();
        $data['lastloginip'] =  request()->ip();
        // $data['lastloginip'] =  ip2long(request()->ip());
     
        // allowField,过滤数组中的非数据表字段数据
        $res = db('users')->where('id', $id)->update(['lastlogintime' => $data['lastlogintime'], 'lastloginip' => $data['lastloginip']]);
        return $res;
    }

}
