import React from "react";
import DialogTitle from "@material-ui/core/DialogTitle";
import DialogContent from "@material-ui/core/DialogContent";
import Typography from "@material-ui/core/Typography";
import TextField from "@material-ui/core/TextField";
import DialogActions from "@material-ui/core/DialogActions";
import {Button} from "@material-ui/core";
import Dialog from "@material-ui/core/Dialog";

class GeneratorLauncher extends React.Component {
    render() {
        return (
            <Dialog onClose={e => {
                this.props.onClose();
            }} aria-labelledby="customized-dialog-title" open={true}>
                <DialogTitle id="customized-dialog-title" onClose={e => {
                }}>
                    Запуск автогенератора
                </DialogTitle>
                <DialogContent dividers>
                    <Typography gutterBottom>
                        Внимание! При запуске генератора будут удалены все старые расписания.
                        Рекомендуется делать это до начала учебного года.
                        Если вы уверены. введите код ниже:<br /><br />
                        <b>666323</b>
                    </Typography>
                    <div>
                        <TextField
                            variant="filled"
                            label={"Код запуска"}
                            x-value={""}
                            disabled={false}
                        />
                    </div>

                </DialogContent>
                <DialogActions>
                    <Button autoFocus disabled={true} onClick={e => {

                    }} color="primary">
                        ПУСК
                    </Button>
                </DialogActions>
            </Dialog>
        );
    }
}

export default GeneratorLauncher;