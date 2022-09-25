import React from 'react';
import {DataGrid} from '@material-ui/data-grid';
import ApiHelper from "../ApiHelper";
import Button from '@material-ui/core/Button';
import {ButtonGroup} from "@material-ui/core";
import PreLoader from "./PreLoader";

class KlassRaumenPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            loading: true,
            items: [],
            selectionModel: []
        };

        setTimeout(() => {
            if (this.state.items.length < 1)
                this.componentDidMount();
        }, 100);
    }

    componentDidMount() {
        ApiHelper.callApiAuthorized(
            '/api/admin/lessons/list',
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
            {field: 'name', headerName: 'Название', width: 150, editable: true},
            {field: 'cabinet', headerName: 'Кабинет', width: 150, editable: true},
        ];

        return <div>
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
                                $fd.append('name', $item.name);
                                $fd.append('cabinet', $item.cabinet);

                                ApiHelper.callApiAuthorized(
                                    '/api/admin/lessons/save',
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
                        "Введите нозвание нового предмета:",
                        "Информатика"
                    );

                    if (name) {
                        let $fd = new FormData();
                        $fd.append('name', name);

                        ApiHelper.callApiAuthorized(
                            '/api/admin/lessons/add',
                            $fd,
                            oData => {
                                this.componentDidMount();
                            }
                        );
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
                            "Вы действительно хотите удалить предметы в количестве " +
                            this.state.selectionModel.length +
                            "?"
                        );
                        if (ret) {
                            let $fd = new FormData();
                            this.state.selectionModel.map(item => {
                                $fd.append('id', item);
                            });

                            ApiHelper.callApiAuthorized(
                                '/api/admin/lessons/delete',
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

export default KlassRaumenPage;