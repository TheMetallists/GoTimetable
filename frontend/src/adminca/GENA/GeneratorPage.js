import React from 'react'
import {Button} from "@material-ui/core";
import GeneratorLauncher from './GeneratorLauncher';
import GeneratorTable from './GeneratorTable';
import ButtonGroup from '@material-ui/core/ButtonGroup';
import ApiHelper from "../../ApiHelper";
import PreLoader from "../PreLoader";

class GeneratorPage extends React.Component {
    constructor() {
        super();
        this.state = {
            isGenaOpen: false,
            loading: true
        };
        this.timeout = false;

        //setInterval(() => this.processChangesIfNecessary(), 5000);
    }

    componentDidMount() {
        this.loadData(-1);
    }

    loadData(parallel) {
        this.setState({
            loading: true,
        }, () => {
            let $fd = new FormData();
            $fd.append('parallel', parallel);

            ApiHelper.callApiAuthorized(
                '/api/admin/gena/list',
                $fd,
                result => {
                    this.setState({
                        loading: false,
                        ...result
                    });
                }
            );
        })

    }

    onAddLesson(l) {
        this.setState({
            loading: true,
        }, () => {
            let $fd = new FormData();
            $fd.append('parallel', this.state.currentParallel.number);
            $fd.append('lesson', l);

            ApiHelper.callApiAuthorized(
                '/api/admin/gena/add',
                $fd,
                result => {
                    this.loadData(this.state.currentParallel.number);
                }
            );
        })
    }

    processChangesIfNecessary() {
        if (this.state.loading)
            return;
        let currentParallel = this.state.currentParallel;
        let hasMods = false;
        Object.keys(currentParallel.elements).map(klass => {
            Object.keys(currentParallel.elements[klass]).map(lesson => {
                let lsn = currentParallel.elements[klass][lesson];
                if (lsn.hasOwnProperty('modded') && lsn.modded) {
                    hasMods = true;

                    let hpw = parseInt(lsn.hoursPerWeek);

                    if (isNaN(hpw))
                        hpw = 1;

                    if (hpw < 0)
                        hpw = parseInt(Math.abs(hpw));

                    currentParallel.elements[klass][lesson].hoursPerWeek = hpw;
                }
            });
        });

        this.setState({
            currentParallel: {...currentParallel},
            loading: hasMods
        }, () => {
            if (!hasMods)
                return;

            let currentParallel = this.state.currentParallel;
            Object.keys(currentParallel.elements).map(klass => {
                Object.keys(currentParallel.elements[klass]).map(lesson => {
                    let lsn = currentParallel.elements[klass][lesson];
                    if (lsn.hasOwnProperty('modded') && lsn.modded) {
                        currentParallel.elements[klass][lesson].modded = false;
                        let hpw = parseInt(lsn.hoursPerWeek);

                        let $fd = new FormData();
                        $fd.append('id', lsn.id);
                        $fd.append('hours', hpw);

                        ApiHelper.callApiAuthorized(
                            '/api/admin/gena/save',
                            $fd,
                            result => {
                                this.setState({
                                    loading: false
                                });
                                //this.loadData(this.state.currentParallel.number);
                            }
                        );

                        currentParallel.elements[klass][lesson].hoursPerWeek = hpw;
                    }
                });
            });

            this.setState({
                currentParallel: {...currentParallel}
            }, () => {
                this.loadData(this.state.currentParallel.number);
            });


        });
    }

    render() {
        if (this.state.loading) {
            return <PreLoader/>
        }

        console.log('STATA: ', this.state);

        return (
            <div className="generator">
                <h1>
                    Генератор расписаний
                </h1>

                <div className="ret-baton-container">
                    <Button
                        color={"secondary"}
                        variant="contained"
                        onClick={e => {
                            this.setState({
                                isGenaOpen: true
                            });
                        }}>
                        Сгенерировать
                    </Button>
                </div>


                <GeneratorTable
                    klassen={this.state.currentParallel.klasses}
                    lessons={this.state.lessons}
                    config={this.state.currentParallel.elements}
                    onAddLesson={l => this.onAddLesson(l)}
                    onSetTime={(klass, lesson, hours) => {
                        let currentParallel = this.state.currentParallel;
                        currentParallel.elements[klass][lesson].hoursPerWeek = hours;
                        currentParallel.elements[klass][lesson].modded = true;
                        this.setState({
                            currentParallel: {...currentParallel}
                        }, () => {
                            if (this.timeout)
                                clearTimeout(this.timeout);
                            this.timeout = setTimeout(() => {
                                this.timeout = false;
                                this.processChangesIfNecessary();
                            }, 5000);
                        });
                    }}
                />

                Выберите параллель:
                <ButtonGroup
                    value={1}
                    exclusive
                    onChange={e => {
                    }}
                    aria-label="text alignment"
                >
                    {this.state.parallels.map(kls =>
                        <Button
                            key={kls}
                            color={(kls == this.state.currentParallel.number) ? "primary" : null}
                            onClick={e => {
                                this.loadData(kls);
                            }}
                        >
                            {kls}
                        </Button>
                    )}
                </ButtonGroup>


                {this.state.isGenaOpen ?
                    <GeneratorLauncher onClose={() => {
                        this.setState({
                            isGenaOpen: false
                        });
                    }}/>
                    : null}
            </div>
        );
    }
}

export default GeneratorPage;