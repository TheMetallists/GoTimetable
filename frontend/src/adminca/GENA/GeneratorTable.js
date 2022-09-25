import React from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import {withStyles} from '@material-ui/core/styles';
import {Button, MenuItem, Select} from "@material-ui/core";
import TextField from "@material-ui/core/TextField";

const StyledTableCell = withStyles((theme) => ({
    head: {
        backgroundColor: '#3f51b5',
        color: theme.palette.common.white,
    },
    body: {
        fontSize: 14,
    },
}))(TableCell);

const StyledTableRow = withStyles((theme) => ({
    root: {
        '&:nth-of-type(odd)': {
            backgroundColor: theme.palette.action.hover,
        },
    },
}))(TableRow);


class GeneratorTable extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            addLessn: (this.props.lessons.length > 0) ? this.props.lessons[0].id : 0
        }
    }

    render() {

        return (
            <TableContainer className="geno-table" component={Paper}>
                <Table>
                    <TableHead>
                        <StyledTableRow>
                            <StyledTableCell>Класс/предмет</StyledTableCell>
                            {this.props.klassen.map(klass => (
                                <StyledTableCell key={klass.id}>
                                    {klass.letter} &nbsp;
                                    ({(() => {
                                    let hrs = 0;
                                    if (this.props.config.hasOwnProperty(klass.id))
                                        Object.values(this.props.config[klass.id]).map(lesson => {
                                            hrs += parseInt(lesson.hoursPerWeek);
                                        });
                                    return hrs;
                                })()} ч)
                                </StyledTableCell>
                            ))}
                        </StyledTableRow>
                    </TableHead>
                    <TableBody>
                        {this.props.lessons.filter(lesson => {
                            return lesson.used;
                        }).map(lesson => {
                            return (
                                <StyledTableRow>
                                    <TableCell>
                                        {lesson.name}
                                    </TableCell>
                                    {this.props.klassen.map(klass => {
                                        return (
                                            <TableCell style={{
                                                background: this.props.config[klass.id][lesson.id].modded ? 'yellow' : null
                                            }}>
                                                <TextField
                                                    value={this.props.config[klass.id][lesson.id].hoursPerWeek}
                                                    type="number"
                                                    onChange={e => {
                                                        this.props.onSetTime(
                                                            klass.id,
                                                            lesson.id,
                                                            e.target.value
                                                        );
                                                    }}
                                                />
                                            </TableCell>
                                        );
                                    })}
                                </StyledTableRow>
                            );
                        })}

                        <StyledTableRow>
                            <TableCell>

                                <div style={{
                                    display: 'flex',
                                    justifyContent: 'space-between'
                                }}>
                                    <Select
                                        value={this.state.addLessn}
                                        onChange={e => {
                                            this.setState({
                                                addLessn: e.target.value
                                            })
                                        }}
                                    >
                                        {this.props.lessons.filter(item => !item.used).map(lesson => (
                                            <MenuItem key={lesson.id} value={lesson.id}>
                                                {lesson.name}
                                            </MenuItem>
                                        ))}
                                    </Select>

                                    <Button
                                        variant="contained"
                                        onClick={e => {
                                            this.props.onAddLesson(this.state.addLessn)
                                        }}
                                    >
                                        Добавить предмет
                                    </Button>
                                </div>

                            </TableCell>
                        </StyledTableRow>
                    </TableBody>
                </Table>
            </TableContainer>
        );
    }
}

export default GeneratorTable;