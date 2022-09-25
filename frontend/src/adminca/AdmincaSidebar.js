import React from 'react';
import Drawer from '@material-ui/core/Drawer';
import List from '@material-ui/core/List';
import Divider from '@material-ui/core/Divider';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import FolderIcon from '@material-ui/icons/Folder';
import SpeakerIcon from '@material-ui/icons/Speaker';
import MenuBookIcon from '@material-ui/icons/MenuBook';
import CastForEducationIcon from '@material-ui/icons/CastForEducation';
import SettingsIcon from '@material-ui/icons/Settings';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import {Redirect} from 'react-router-dom';

class AdmincaSidebar extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            redirect: false
        };
    }

    render() {
        if (this.state.redirect) {
            setTimeout(() => {
                this.setState({
                    redirect: false
                });
            }, 200);
            return (
                <Redirect to={this.state.redirect}/>
            );
        }
        let $linkMap = [
            {name: 'Кабинеты (предметы)', url: '/adminca/klassrummen/', icon: <MenuBookIcon/>},
            {name: 'Классы', url: '/adminca/klassen/', icon: <CastForEducationIcon/>},
            {name: 'Генератор расписаний',url:'/adminca/generaten/',icon:<SettingsIcon/>},
            {name: 'Матюгальник', url: '/adminca/boombox/', icon: <SpeakerIcon/>},
            {name: 'Филе', url: '/adminca/file/', icon: <FolderIcon/>},
        ];
        let $hasHeading = false;
        let $openKlassMode = null;
        let $kmc = document.location.pathname.match(/\/adminca\/klassen\/([0-9]+)\//);
        if ($kmc) {
            $openKlassMode = 'Класс #' + $kmc[1];
            $hasHeading = true;
        }

        return [(
            <AppBar position="fixed" className={this.props.tbClass}>
                <Toolbar>
                    <Typography variant="h6" noWrap>
                        {Object.values($linkMap).map($arItem => {
                            if (window.location.pathname === $arItem.url) {
                                $hasHeading = true;
                                return $arItem.name;
                            }
                        })}
                        {$openKlassMode}
                        {$hasHeading ? null : 'Админка'}
                    </Typography>
                </Toolbar>
            </AppBar>
        ), (
            <Drawer
                variant="permanent"
                anchor="left"
                className={this.props.className}
                classes={this.props.classes}
            >
                <div/>
                <Divider/>
                <List>
                    {Object.values($linkMap).map(($arItem,b) => (
                        <ListItem
                            button
                            key={b}
                            onClick={e => {
                                this.setState({
                                    redirect: $arItem.url
                                })
                            }}
                        >
                            <ListItemIcon>{$arItem.icon}</ListItemIcon>
                            <ListItemText primary={$arItem.name}/>
                        </ListItem>
                    ))}
                </List>
                {/*<Divider/>
                <List>
                    {['All mail', 'Trash', 'Spam'].map((text, index) => (
                        <ListItem button key={text}>
                            <ListItemIcon>{index % 2 === 0 ? <InboxIcon/> : <MailIcon/>}</ListItemIcon>
                            <ListItemText primary={text}/>
                        </ListItem>
                    ))}
                </List>*/}
            </Drawer>
        )];
    }
}

export default AdmincaSidebar;