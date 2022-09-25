import React, {Suspense} from "react";
import AuthorizationForm from './AuthorizationForm';
import AdmincaSidebar from './AdmincaSidebar';
import Toolbar from '@material-ui/core/Toolbar';
import {withStyles} from "@material-ui/core/styles";
import {Redirect, Route, Switch} from 'react-router-dom';
import ApiHelper from "../ApiHelper";
import PreLoader from './PreLoader';

const KlassRaumenPage = React.lazy(() => import('./KlassRaumenPage'));
const KlassenPage = React.lazy(() => import('./KlassenPage'));
const KlassenLessnsPage = React.lazy(() => import('./KlassenLessnsPage'));
const GeneratorPage = React.lazy(() => import('./GENA/GeneratorPage'));

const drawerWidth = 240;

const styles = (theme) => ({
    root: {
        display: 'flex',
    },
    appBar: {
        width: `calc(100% - ${drawerWidth}px)`,
        marginLeft: drawerWidth,
    },
    drawer: {
        width: drawerWidth,
        flexShrink: 0,
    },
    drawerPaper: {
        width: drawerWidth,
    },
    // necessary for content to be below app bar
    toolbar: theme.mixins.toolbar,
    content: {
        flexGrow: 1,
        backgroundColor: theme.palette.background.default,
        padding: theme.spacing(3),
    },
});

class AdmincaMain extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            authorized: false,
            authToken: ''
        };
    }

    checkAndPushStoredToken() {
        if (typeof window.localStorage === 'undefined')
            return;

        const $token = window.localStorage.getItem('auth.token', 'false');
        if (!$token || $token === 'false' || $token === 'null') {
            //alert($token);
            return;
        }

        ApiHelper.token = $token;
        ApiHelper.callApiAuthorized('/api/admin/test', null, oItem => {
            //
        }, e => {
            this.setState({
                authToken: $token,
                authorized: true
            });
        });

    }

    componentDidMount() {
        this.checkAndPushStoredToken();
    }

    render() {
        if (!this.state.authorized) {
            return (
                <AuthorizationForm
                    onAuthorized={authToken => {
                        if (typeof window.localStorage !== 'undefined') {
                            window.localStorage.setItem('auth.token', authToken)
                        }
                        this.setState({
                            authToken,
                            authorized: true
                        });
                    }}
                />
            );
        }

        ApiHelper.token = this.state.authToken;

        const {classes} = this.props;
        return (
            <div className={classes.root}>
                <AdmincaSidebar
                    className={classes.drawer}
                    classes={{
                        paper: classes.drawerPaper,
                    }}
                    tbClass={classes.appBar}
                />

                <main className={classes.content}>
                    <Toolbar/>
                    <Switch>
                        <Route path={"/adminca/klassrummen/"}>
                            <Suspense fallback={<PreLoader/>}>
                                <KlassRaumenPage
                                    authToken={this.state.authToken}
                                />
                            </Suspense>
                        </Route>
                        <Route path={"/adminca/klassen/"} exact>
                            <Suspense fallback={<PreLoader/>}>
                                <KlassenPage
                                    authToken={this.state.authToken}
                                />
                            </Suspense>
                        </Route>

                        <Route path={'/adminca/klassen/:klid/'}>
                            <Suspense fallback={<PreLoader/>}>
                                <KlassenLessnsPage
                                    authToken={this.state.authToken}
                                />
                            </Suspense>
                        </Route>

                        <Route path={"/adminca/generaten/"} exact>
                            <Suspense fallback={<PreLoader />}>
                                <GeneratorPage
                                    authToken={this.state.authToken}
                                    />
                            </Suspense>
                        </Route>

                        <Route path={"/adminca/boombox/"}>
                            <div>asdasdasdqwe</div>
                            <PreLoader/>
                        </Route>

                        <Route path={"/adminca/file/"}>
                            <div>asdasdasdqwe</div>
                        </Route>
                    </Switch>
                </main>

            </div>
        );
    }
}

export default withStyles(styles, {withTheme: true})(AdmincaMain);