import React, {Suspense} from 'react';
import {BrowserRouter as Router, Redirect, Route, Switch} from 'react-router-dom';

const TTableContainer = React.lazy(() => import('./telek/TTableContainer'));
const AdmincaMain = React.lazy(() => import('./adminca/AdmincaMain'));

class RouterObject extends React.Component {
    render() {
        return (
            <Router>
                <Switch>
                    <Route path="/" exact>
                        <Redirect to={"/telek"}/>
                    </Route>

                    <Route path={"/telek"} exact>
                        <Suspense fallback={<div>Ждите...</div>}>
                            <TTableContainer/>
                        </Suspense>
                    </Route>

                    <Route path={"/adminca"}>
                        <Suspense fallback={<div>Ждите...</div>}>
                            <AdmincaMain/>
                        </Suspense>
                    </Route>

                    <Route path='*' exact={true}>
                        <h1>404 Not Found</h1><br/>
                        <i>Watcha lookin' 4?</i>
                    </Route>
                </Switch>
            </Router>
        );
    }
}

export default RouterObject;