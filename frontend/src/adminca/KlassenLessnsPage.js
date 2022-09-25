import React from 'react';
import {Redirect, withRouter} from 'react-router-dom';
import ApiHelper from "../ApiHelper";
import PreLoader from "./PreLoader";

import WeekDayItem from './WeekDayItem';
import {Button} from "@material-ui/core";

class KlassenLessnsPage extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            loading: true,
            is404: false,
            items: {}
        };

        setTimeout(() => {
            if (Object.values(this.state.items).length < 1)
                this.componentDidMount();
        }, 200);
    }

    componentDidMount() {
        let $df = new FormData();
        $df.append('id', this.props.match.params.klid);
        ApiHelper.callApiAuthorized(
            '/api/admin/timetables/list',
            $df,
            $oData => {
                this.setState({
                    loading: false,
                    ...$oData
                });
            }, e => {
                this.setState({
                    is404: true
                });
            }
        );
    }

    doSaveDataset() {
        ApiHelper.callApiAuthorizedJSON(
            '/api/admin/timetables/commit',
            JSON.stringify({
                kid: parseInt(this.props.match.params.klid),
                items: this.state.items
            }),
            oResult => {
                alert('Сохранено!');
                this.setState({
                    loading: true
                }, () => this.componentDidMount());
            }
        );
    }

    render() {
        if (this.state.is404) {
            return <Redirect to={'/adminca/klassen/'}/>;
        }

        if (this.state.loading) {
            return <PreLoader/>;
        }


        console.log(this.state);
        return (
            <div>
                <h1 style={{
                    fontSize: '2rem',
                    fontFamily: '"Roboto", "Helvetica", "Arial", sans-serif'
                }}>
                    Расписание класса "{this.state.klass.number}{this.state.klass.letter}"
                </h1>

                <Button
                    color={"primary"}
                    variant="contained"
                    onClick={e => {
                        this.doSaveDataset();
                    }}>
                    Сохранить
                </Button>

                {[1, 2, 3, 4, 5, 6].map(iWeekDay => {
                    return <WeekDayItem
                        key={iWeekDay}
                        {...this.state}
                        iWeekDay={iWeekDay}
                        klid={this.props.match.params.klid}
                        items={this.state.items[iWeekDay]}
                        onSetItems={newItems => {
                            let $items = {...this.state.items};
                            $items[iWeekDay] = newItems;
                            this.setState({
                                items: {...$items}
                            });
                        }}
                    />;
                })}

                <div style={{
                    margin: '25px 0'
                }}>
                    <Button
                        color={"primary"}
                        variant="contained"
                        onClick={e => {
                            this.doSaveDataset();
                        }}>
                        Сохранить
                    </Button>
                </div>

            </div>
        );
    }
}

export default withRouter(KlassenLessnsPage);