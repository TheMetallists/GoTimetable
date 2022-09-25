import React from "react";
import {Button} from '@material-ui/core';
import Dialog from '@material-ui/core/Dialog';
import DialogTitle from '@material-ui/core/DialogTitle';
import DialogContent from '@material-ui/core/DialogContent';
import DialogActions from '@material-ui/core/DialogActions';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import ApiHelper from '../ApiHelper.js';

class AuthorizationForm extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            pass: ''
        };
    }

    authorize() {
        let $fd = new FormData();
        $fd.append("password", this.state.pass);
        ApiHelper.callApiUnauthorized("/api/auth", $fd, oData => {
            this.props.onAuthorized(oData.token);
        });
    }

    render() {
        return (
            <Dialog onClose={e => {
                alert("Вы не вошли в систему, чтобы что-то закрывать!")
            }} aria-labelledby="customized-dialog-title" open={true}>
                <DialogTitle id="customized-dialog-title" onClose={e => {
                }}>
                    Авторизация
                </DialogTitle>
                <DialogContent dividers>
                    <Typography gutterBottom>
                        Для входа введите логин и пароль.
                    </Typography>
                    <div>
                        <TextField
                            variant="filled"
                            label={"Login"}
                            value={"admin"}
                            disabled={true}
                        />
                        <TextField
                            variant="filled"
                            type="password"
                            label={"Password"}
                            value={this.state.pass}
                            onChange={e => {
                                this.setState({
                                    pass: e.target.value
                                });
                            }}
                        />
                    </div>

                </DialogContent>
                <DialogActions>
                    <Button autoFocus onClick={e => {
                        this.authorize();
                    }} color="primary">
                        Войти
                    </Button>
                </DialogActions>
            </Dialog>
        );
    }
}

export default AuthorizationForm;