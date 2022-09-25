import React from 'react';
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';
import {DataGrid} from '@material-ui/data-grid';
import {Button, ButtonGroup} from "@material-ui/core";
import ApiHelper from "../ApiHelper";

class WeekDayItem extends React.Component {
    constructor() {
        super();
        this.state = {
            selectionModel: []
        };
    }

    render() {
        if (Object.keys(this.props.lessons).length < 1)
            return <div>Добавьте хотя бы 1 предмет!</div>;

        const columns = [
            {field: 'id', headerName: '#', width: 90},
            {field: 'number', headerName: 'Номер', width: 150, editable: false},
            {
                field: "lesson",
                headerName: "Урок",
                width: 200,
                disableClickEventBubbling: true,
                renderCell: (params) => {
                    //console.log('PARAMS', params);

                    return (
                        <Select
                            value={params.value}
                            onChange={e => {
                                let $items = [...this.props.items];

                                for (let idy = 0; idy < $items.length; idy++) {
                                    if ($items[idy].id === params.id) {
                                        $items[idy].lesson = e.target.value;
                                        break;
                                    }
                                }

                                this.props.onSetItems($items);
                            }}
                        >
                            {Object.values(this.props.lessons).map($arItem => {
                                return (
                                    <MenuItem key={$arItem.id} value={$arItem.id}>
                                        {$arItem.name} (каб. {$arItem.cabinet})
                                    </MenuItem>
                                );
                            })}
                        </Select>
                    );
                }
            },
        ];

        const $wdys = [
            'Будущее',
            'Понедельник',
            'Вторник',
            'Среда',
            'Четверг',
            'Пятница',
            'Суббота',
        ];
        const {iWeekDay} = this.props;

        let lnum = 1;
        return (
            <div key={iWeekDay}>
                <h3 style={{
                    fontFamily: '"Roboto", "Helvetica", "Arial", sans-serif'
                }}>
                    {$wdys[iWeekDay]}
                </h3>
                <div style={{height: 400, width: '100%'}}>
                    {(this.props.items && this.props.items.length > 0) ? (
                        <DataGrid
                            rows={this.props.items.map($arItem => ({
                                id: $arItem.id,
                                number: lnum++,
                                lesson: $arItem.lesson
                            }))}
                            columns={columns}
                            pageSize={5}
                            checkboxSelection
                            onSelectionModelChange={(newSelection) => {
                                console.log(newSelection.selectionModel);
                                this.setState({
                                    selectionModel: newSelection.selectionModel
                                });
                            }}
                            selectionModel={this.state.selectionModel}
                        />
                    ) : null}
                </div>
                <ButtonGroup color="primary" aria-label="">
                    <Button onClick={e => {
                        // DataGrid requires a unique and positive ID. There's only one way to get it correctly.
                        let $fd = new FormData();
                        $fd.append('klass', this.props.klid);
                        $fd.append('lesson', parseInt(Object.keys(this.props.lessons)[0]));
                        $fd.append('weekday', this.props.iWeekDay);

                        ApiHelper.callApiAuthorized(
                            '/api/admin/timetables/add', $fd, oData => {
                                if (oData.id > 0) {
                                    let $items = [];
                                    if (typeof this.props.items !== 'undefined')
                                        $items = [...this.props.items];
                                    $items.push({
                                        id: oData.id,
                                        weekday: this.props.iWeekDay,
                                        lesson: parseInt(Object.keys(this.props.lessons)[0])
                                    });

                                    this.props.onSetItems($items);
                                }
                            }
                        );

                    }}>
                        Добавить
                    </Button>
                    <Button
                        color={"secondary"}
                        disabled={this.state.selectionModel.length < 1}
                        onClick={e => {
                            //TODO: add a nice popup
                            var ret = window.confirm(
                                "Вы действительно хотите удалить урок?"
                            );
                            if (ret) {
                                let $items = [...this.props.items];
                                $items = $items.filter(arItem => {
                                    return !(this.state.selectionModel.indexOf(arItem.id) > -1);
                                })
                                this.props.onSetItems($items);
                            }
                        }}>
                        Удалить
                    </Button>
                </ButtonGroup>
            </div>
        );
    }
}

export default WeekDayItem;