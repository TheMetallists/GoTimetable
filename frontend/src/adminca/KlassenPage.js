import React from "react";
import {DataGrid} from '@material-ui/data-grid';
import {Button, ButtonGroup} from "@material-ui/core";
import ApiHelper from "../ApiHelper";
import {Redirect} from "react-router-dom";
import PreLoader from "./PreLoader";

class KlassenPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            items: [],
            selectionModel: [],
            raspOpen: false,
            loading: true
        };

        setTimeout(() => {
            if (this.state.items.length < 1)
                this.componentDidMount();
        }, 100);
    }

    componentDidMount() {
        ApiHelper.callApiAuthorized(
            '/api/admin/klassen/list',
            new FormData(),
            oData => {
                this.setState({
                    items: oData.items,
                    loading: false
                })
            }
        );
    }

    render() {
        if (this.state.loading)
            return <PreLoader/>;
        const columns = [
            {field: 'id', headerName: '#', width: 90},
            {field: 'number', headerName: 'Параллель', width: 150, editable: true},
            {field: 'letter', headerName: 'Буква', width: 150, editable: true},
            {field: 'countLessons', headerName: 'Уроков/нед.', width: 150, editable: false},
            {
                field: "",
                headerName: "Action",
                width: 200,
                disableClickEventBubbling: true,
                renderCell: (params) => {
                    const onClick = e => {
                        console.log('CLK: ', params.row.id);
                        this.setState({
                            raspOpen: params.row.id
                        });
                    };

                    return <Button variant={"contained"} onClick={onClick}>
                        Расписание
                    </Button>;
                }
            },
        ];

        return <div>
            {this.state.raspOpen ? (
                <Redirect to={"/adminca/klassen/" + this.state.raspOpen + "/"}/>
            ) : null}
            <div style={{height: 400, width: '100%'}}>
                <DataGrid
                    rows={this.state.items}
                    columns={columns}
                    pageSize={50}
                    checkboxSelection
                    onSelectionModelChange={(newSelection) => {
                        console.log(newSelection.selectionModel);
                        this.setState({
                            selectionModel: newSelection.selectionModel
                        });
                    }}
                    selectionModel={this.state.selectionModel}
                    onEditCellChangeCommitted={(event) => {
                        this.state.items.map($arItem => {
                            if ($arItem.id === event.id) {
                                let $item = $arItem;
                                $item[event.field] = event.props.value;

                                let $fd = new FormData();
                                $fd.append('id', $item.id);
                                $fd.append('letter', $item.letter);
                                $fd.append('number', $item.number);

                                ApiHelper.callApiAuthorized(
                                    '/api/admin/klassen/save',
                                    $fd,
                                    oData => {
                                        this.componentDidMount();
                                    }
                                );

                            }
                        })
                    }}
                />
            </div>
            <ButtonGroup color="primary" aria-label="">
                <Button onClick={e => {
                    //TODO: add a nice popup
                    var name = window.prompt(
                        "Введите нозвание нового класса:",
                        "3В"
                    );

                    if (!name)
                        return;

                    const $match = name.match(/^([0-9]{1,2})([А-Я])$/);
                    if ($match) {
                        let $fd = new FormData();
                        $fd.append('number', $match[1]);
                        $fd.append('letter', $match[2]);

                        ApiHelper.callApiAuthorized(
                            '/api/admin/klassen/add',
                            $fd,
                            oData => {
                                this.componentDidMount();
                            }
                        );
                    } else {
                        //
                    }
                }}>
                    Добавить
                </Button>
                <Button
                    color={"secondary"}
                    disabled={this.state.selectionModel.length < 1}
                    onClick={e => {
                        //TODO: add a nice popup
                        var ret = window.confirm(
                            "Вы действительно хотите удалить классы и их уроки в количестве " +
                            this.state.selectionModel.length +
                            "?"
                        );
                        if (ret) {
                            let $fd = new FormData();
                            this.state.selectionModel.map(item => {
                                $fd.append('id', item);
                            });

                            ApiHelper.callApiAuthorized(
                                '/api/admin/klassen/delete',
                                $fd,
                                oData => {
                                    this.setState({
                                        selectionModel: []
                                    }, () => this.componentDidMount());
                                }
                            );
                        }
                    }}>
                    Удалить
                </Button>
            </ButtonGroup>
        </div>;
    }

}

export default KlassenPage;