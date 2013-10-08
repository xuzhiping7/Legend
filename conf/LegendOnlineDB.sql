-- phpMyAdmin SQL Dump
-- version 3.4.10.1
-- http://www.phpmyadmin.net
--
-- 主机: localhost
-- 生成日期: 2013 年 10 月 08 日 15:18
-- 服务器版本: 5.5.20
-- PHP 版本: 5.3.10

SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- 数据库: `studygolang`
--

-- --------------------------------------------------------

--
-- 表的结构 `wechat_base`
--

CREATE TABLE IF NOT EXISTS `wechat_base` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '索引ID',
  `openid` varchar(40) NOT NULL COMMENT '微信OPENID',
  `nickname` varchar(20) NOT NULL COMMENT '用户昵称--角色名',
  `username` varchar(20) NOT NULL COMMENT '用户帐户名',
  `sex` tinyint(2) NOT NULL DEFAULT '0' COMMENT '性别：0-未选择，1-男，2-女，3-其他',
  `level` int(10) NOT NULL DEFAULT '0' COMMENT '等级',
  `exp` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '当前等级所在经验',
  `coin` int(11) NOT NULL DEFAULT '0' COMMENT '金币数量',
  `mobility` int(10) unsigned NOT NULL DEFAULT '100' COMMENT '行动力',
  `reputation` int(10) NOT NULL DEFAULT '0' COMMENT '声望值',
  `attack` smallint(5) NOT NULL DEFAULT '0' COMMENT '攻击力点数',
  `defense` smallint(5) NOT NULL DEFAULT '0' COMMENT '防御力点数',
  `stamina` smallint(5) NOT NULL DEFAULT '0' COMMENT '耐力点数',
  `agility` smallint(5) NOT NULL DEFAULT '0' COMMENT '敏捷力点数',
  `wisdom` smallint(5) NOT NULL DEFAULT '0' COMMENT '智慧值',
  `no_distribution` smallint(5) NOT NULL DEFAULT '0' COMMENT '未分配点数',
  `location` mediumint(9) NOT NULL DEFAULT '0' COMMENT '玩家所在位置',
  `flag` int(11) NOT NULL DEFAULT '0' COMMENT '角色主线FLAG进度',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='mobility' AUTO_INCREMENT=5 ;

--
-- 转存表中的数据 `wechat_base`
--

INSERT INTO `wechat_base` (`id`, `openid`, `nickname`, `username`, `sex`, `level`, `exp`, `coin`, `mobility`, `reputation`, `attack`, `defense`, `stamina`, `agility`, `wisdom`, `no_distribution`, `location`, `flag`) VALUES
(1, 'xuzhipingtest', '随风', '好好喝', 0, 10, 100, 3500, 94, 23, 5, 5, 5, 5, 5, 30, 1, 3),
(2, 'xuzhipingtest2', '德玛西亚王子', 'EmptyNow', 0, 0, 2, 5000, 100, 0, 5, 5, 5, 5, 5, 0, 1, 3),
(3, 'aaa', 'EmptyNow', 'EmptyNow', 0, 0, 0, 0, 100, 0, 5, 5, 5, 5, 5, 0, 0, 0),
(4, 'gdsgdsg', '德玛西亚', 'EmptyNow', 0, 0, 0, 0, 100, 0, 5, 5, 5, 5, 5, 0, 0, 3);

-- --------------------------------------------------------

--
-- 表的结构 `wechat_language_template`
--

CREATE TABLE IF NOT EXISTS `wechat_language_template` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `number` int(10) NOT NULL COMMENT '序号',
  `template` varchar(1024) CHARACTER SET utf8 NOT NULL COMMENT '模板',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=77 ;

--
-- 转存表中的数据 `wechat_language_template`
--

INSERT INTO `wechat_language_template` (`id`, `number`, `template`) VALUES
(2, 1, '请输出‘我’查看您的最新状态。输入‘帮助’可以查看相关操作指令。'),
(3, 2, '注册'),
(4, 3, '欢迎来到微信奇幻网游《传说》，请输出''注册''，确认注册游戏。'),
(5, 4, '创建角色中,请输入您的角色名。(例如‘柳随风’，2~8个汉字内。)'),
(6, 5, '角色【%s】命名成功！你可以输入''帮助''查看简单的游戏指导。'),
(7, 6, '%s\n当前地点:%s\n等级:%d  经验:%d/%d\n行动力:%d/%d\n声望:%d\n金币:%d\n\n生命:%d/%d\n负重:%d/%d\n抗性:%d\n\n攻击:%d   防御:%d\n体力:%d   敏捷:%d\n智慧:%d\n剩余分配点数:%d\n\n(您可以输入''帮助''查看更多操作)'),
(8, 7, '抱歉，你不能到达这个地方。请输入''前往''查看你当前能前往的地方。'),
(9, 8, '你到达了%s'),
(10, 9, '前往%s\n'),
(11, 10, '你当前可以前往的地方:\n'),
(12, 11, '您可以输入相应关键字来实现的操作：\n【我】\n[查看角色状态]\n【道具】\n[查看已拥有道具]\n【使用】+【道具名】\n[使用相应道具] \n【前往】\n[查看你可以去的地方]\n【前往】+【地名】\n[命令角色前往目标地方]\n【修炼】\n[在角色当前的地方修炼]\n【休息】\n[进入休息状态，每分钟回复1生命，行动性操作会打断此状态。]\n【搜寻】\n[搜寻当前地方的物品或事件，暂无，未完成]\n【附近】\n[查看与你同地图的玩家]\n【挑战】+【角色名】\n[挑战在你附近的角色]\n【提升】\n[分配属性点]\n【事件】\n[最近你发生的重要事件]\n【传说】\n[查看在线人数版本信息]\n'),
(13, 12, '很遗憾，你没有搜寻到你所有用的信息或物品。'),
(14, 100, '网络错误，请重新输入。'),
(15, 10001, '测试测试测试测试'),
(16, 100001, '[提升%s%d]恭喜!你的%s由%d提升到了%d，输入''我''查看当前最新状态。'),
(17, 100002, '没有可用属性分配点，每个等级的提升都能获得额外3点属性分配点，输入‘修炼’进入战斗胜利可获取经验升级。'),
(18, 100003, '[提升]你当前有%d分配点，可用于提升以下属性，输入：\n[攻击]+[数字n]:提升n点攻击\n[防御]+[数字n]:提升n点防御\n[体力]+[数字n]:提升n点体力\n[敏捷]+[数字n]:提升n点敏捷\n[智慧]+[数字n]:提升n点智慧\n(例如：攻击1)\n'),
(19, 100004, '你当前只有[%d]属性分配点，不能提升%d点%s，输入''提升''查看详细。'),
(20, 100005, '背包：'),
(21, 100006, '您当前没有任何道具。''修炼''打怪有一定几率可以获得道具，可以多历练下。'),
(22, 100007, '【%s】× %d ：[%s]'),
(23, 100008, '使用了【%s】,回复生命值%d,当前生命值%d/%d。'),
(24, 100009, '你没有【%s】，不能使用，请查看是否输入错误。'),
(25, 100010, '这个道具似乎没有一丁点用。'),
(26, 100011, '进入休息状态，任何行动性操作将会打断休息。休息状态下，生命值每分钟回复1点，行动力回复1点。'),
(27, 100012, '你当前已经正处于休息状态，输入‘我’可以查看当前状态。'),
(28, 100013, '抱歉，你当前所在的地方没有东西购买。\n请‘前往林风村’尝试购买。输入''前往''能查看当前你所可以去的地方。'),
(29, 100014, '抱歉，你当前所在的地方不能出售物品。\n请‘前往林风村’尝试出售。输入''前往''能查看当前你所可以去的地方。'),
(30, 100015, '你可以购买以下物品：'),
(31, 100016, '【%s】×%d\n[%s 价钱:%d金钱/件]'),
(32, 100017, '抱歉，你所在的地方没有东西购买。'),
(33, 100018, '抱歉，你身上并没有任何东西可以出售。'),
(34, 100019, '你可以出售以下物品：\n'),
(35, 100020, '抱歉，你身上并没有【%s】可用于出售。'),
(36, 100021, '你出售了一件【%s】。\n\n金币增加了%d。目前金钱为%d。\n\n输入''我''可查看最新状态。'),
(37, 100022, '抱歉，你身上只有%d金币，【%s】需要花费%d金币，请到某地找找生财之道再购买此物品。'),
(38, 100023, '你花费%d金币购买了【%s】，输入''道具''可查看所拥有的物品。'),
(39, 100024, '【%s】\n[%s 价钱:%d金币/件]'),
(40, 100025, '抱歉，该村落并没有【%s】可以购买。'),
(41, 100026, '使用了【%s】,回复行动力%d,当前行动力%d/%d。'),
(42, 100027, '你的所有属性提升点已经重置。请输入''我''查看。'),
(43, 100028, '抱歉，此区域比较危险，需要准入等级%d，输入''前往''查看你所能到达的地方。'),
(44, 100029, '请输入新角色名。(例如‘随风’，2~8个汉字内。)'),
(45, 200001, '前往%s'),
(46, 200002, '在%s修炼'),
(47, 200003, '在%s买了%s'),
(48, 200004, '在%s卖出了%s'),
(49, 200005, '进入休息状态。'),
(50, 200006, '使用了道具%s'),
(51, 200007, '暂无任何事件记录。'),
(52, 200008, '%s\n\n%s'),
(53, 800000, '你在%s遇到了%s，你轻而易举地取得了胜利。'),
(54, 800001, '恭喜你升级了！增加3点自由分配属性。输入''提升''查看。'),
(55, 800002, '您的行动力不足以完成这样的事情。\n\n某些操作需要消耗行动力，行动力每天更新，可以输入‘我’查看。'),
(56, 800003, '行动力：%d\n当前行动力：%d/%d'),
(57, 800004, '你在%s遇到了%s，你与它艰难地搏斗了一番，遗憾的是最后体力不支倒下，被附近巡逻队发现，营救回%s。'),
(58, 800005, '生命值：%d\n当前生命值：%d/%d'),
(59, 800006, '经验：+%d'),
(60, 800007, '获得：【%s】[%s]'),
(61, 800008, '抱歉，你当前所在的地方并不适合修炼。\n请‘前往林风南海岸’进行修炼。输入''前往''能查看当前你所可以去的地方。'),
(62, 800009, '到达了%s\n\n%s'),
(63, 800010, '你周围的玩家有：'),
(64, 800011, '你周围并没有其他玩家。请''前往''别的地方寻找一下。'),
(65, 800012, '%s里，没有名叫%s的玩家，请输入''附近''查看下在你附近的人。'),
(66, 800013, '你与%s大战了数十回合，获得了胜利。%s不甘地倒在地上。'),
(67, 800014, '你与%s大战了数十回合，无奈%s技艺略胜一筹，你体力不支倒在地上，等你醒过来时，发现自己已经在%s了。'),
(68, 800015, '你不能挑战自己，请输入''附近''查看下在你附近的人。'),
(69, 800016, '%s向你发起挑战，你们大战了数个回合便分出胜负，遗憾的是%s略胜一筹，你体力不支倒下来，醒来的当时发现自己已经在%s了。'),
(70, 800017, '%s向你发起挑战，你们大战了数个回合便分出胜负，你轻松地战胜了%s。损失了体力%d点。'),
(71, 900000, '版本：Version_0.10813\n当前在线人数:%d人\n\n建议或BUG请反馈:xuzhiping7@qq.com\n微博@葱烧烙饼'),
(72, 900001, '指令错误，输入‘帮助’可以查看相关操作指令。'),
(73, 900002, '系统内部错误，请尝试输入别的命令。'),
(74, 900003, '此命令无法执行，请查看是否错误输入。'),
(75, 0, '恢复得好烦得很喝点水');

-- --------------------------------------------------------

--
-- 表的结构 `wechat_map`
--

CREATE TABLE IF NOT EXISTS `wechat_map` (
  `id` int(8) NOT NULL AUTO_INCREMENT,
  `number` int(8) NOT NULL DEFAULT '0' COMMENT '序号编号',
  `level` int(8) NOT NULL DEFAULT '0' COMMENT '准入等级',
  `name` varchar(30) NOT NULL DEFAULT '0' COMMENT '地图名',
  `descript` varchar(1024) NOT NULL DEFAULT '0' COMMENT '地图描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=10 ;

--
-- 转存表中的数据 `wechat_map`
--

INSERT INTO `wechat_map` (`id`, `number`, `level`, `name`, `descript`) VALUES
(4, 1, 0, '林风村', '林风角的一个小村落，常年大风，在海角之上临海而拔起，空气潮湿。春秋之季常遭海怪袭击。'),
(5, 2, 0, '林风南海岸', '海怪的聚集地，是初出勇者作为''修炼''的好地方。'),
(7, 3, 5, '流放平原', '多年来林城帝国流放重犯之地，妖魂遍野，也有不少勇士在此扎根修炼，但有去无回的传闻实在太多了。'),
(8, 4, 10, '无名村', '无从属国家地区的村落，由浪人自发建立而起，环境十分恶劣。目前仅能提供简单的补给。'),
(9, 5, 10, '临界关', '鲜有人知临界关的后边有什么，临界关迷雾重重，越往深处越无法识别方向。传说在临界关修炼获得‘苍茫之境挑战书’，顺利挑战成功则能顺利到达苍茫之境。这里妖物强大，一不留神，可能将会惨绝人寰。');

-- --------------------------------------------------------

--
-- 表的结构 `wechat_map_moster`
--

CREATE TABLE IF NOT EXISTS `wechat_map_moster` (
  `id` int(8) NOT NULL AUTO_INCREMENT,
  `npc_id` int(8) NOT NULL DEFAULT '0' COMMENT 'npc的ID',
  `moster_id` int(8) NOT NULL DEFAULT '0' COMMENT '怪物的ID',
  `rate` tinyint(3) NOT NULL DEFAULT '0' COMMENT '对应怪物出现的概率（%）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='地图和怪物的对应表' AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `wechat_map_npc`
--

CREATE TABLE IF NOT EXISTS `wechat_map_npc` (
  `id` int(8) NOT NULL AUTO_INCREMENT,
  `map_id` int(8) NOT NULL DEFAULT '0' COMMENT '地图ID',
  `npc_id` int(8) NOT NULL DEFAULT '0' COMMENT 'NPC的ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='地图与NPC的对应表' AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `wechat_npc`
--

CREATE TABLE IF NOT EXISTS `wechat_npc` (
  `id` int(8) NOT NULL AUTO_INCREMENT,
  `name` varchar(10) NOT NULL COMMENT 'NPC名字',
  `type` tinyint(2) NOT NULL COMMENT 'NPC类型',
  `desc_simple` varchar(30) NOT NULL COMMENT 'NPC简单描述',
  `desc_detail` varchar(150) NOT NULL COMMENT 'NPC详细描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='NPC表' AUTO_INCREMENT=5 ;

--
-- 转存表中的数据 `wechat_npc`
--

INSERT INTO `wechat_npc` (`id`, `name`, `type`, `desc_simple`, `desc_detail`) VALUES
(2, '村长', 0, '林风村村长', '林风村村长林风村村长林风村村长'),
(3, '王况', 0, '装备买卖商人', '装备买卖商人装备买卖商人');

-- --------------------------------------------------------

--
-- 表的结构 `wechat_npc_conversation`
--

CREATE TABLE IF NOT EXISTS `wechat_npc_conversation` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `npc_id` int(8) NOT NULL DEFAULT '0',
  `conversation` varchar(256) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='NPC对话表' AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `wechat_player_prop`
--

CREATE TABLE IF NOT EXISTS `wechat_player_prop` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '索引主键',
  `player_id` int(11) NOT NULL COMMENT '玩家ID',
  `prop_id` int(11) NOT NULL COMMENT '道具ID',
  `prop_num` int(11) NOT NULL COMMENT '道具数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 COMMENT='玩家道具表' AUTO_INCREMENT=27 ;

--
-- 转存表中的数据 `wechat_player_prop`
--

INSERT INTO `wechat_player_prop` (`id`, `player_id`, `prop_id`, `prop_num`) VALUES
(24, 1, 3, 29),
(25, 1, 1, 33),
(26, 1, 2, 3);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
