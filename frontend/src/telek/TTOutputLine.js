import React from 'react';

class TTOutputLine extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            scrollIndex: 0
        };

        setInterval(() => {
            //this.onScrollRequired();
        }, 5000);
    }

    onScrollRequired() {
        const $maxlen = Object.keys(this.props.timetable).length - 1;
        if (this.state.scrollIndex < $maxlen) {
            this.setState({
                scrollIndex: this.state.scrollIndex + 1
            });
        } else {
            this.setState({
                scrollIndex: 0
            });
        }
    }


    render() {
        let $arFiller = [];
        for (let a = 0; a < 7; a++) {
            $arFiller.push(<div className="classWeekday"/>);
        }
        return (
            <div className="gridtable">
                <div className="classHeader"/>
                <div className="ftoper"/>
                <div className="classWeekday ttoper">Понедельник</div>
                {$arFiller}
                <div className="ftoper"/>
                <div className="classWeekday ttoper">Вторник</div>
                {$arFiller}
                <div className="ftoper"/>
                <div className="classWeekday ttoper">Среда</div>
                {$arFiller}
                <div className="ftoper"/>
                <div className="classWeekday ttoper">Четверг</div>
                {$arFiller}
                <div className="ftoper"/>
                <div className="classWeekday ttoper">Пятница</div>
                {$arFiller}
                <div className="ftoper"/>
                <div className="classWeekday ttoper">Суббота</div>
                {$arFiller}


                {Object.keys(this.props.timetable).slice(this.state.scrollIndex, this.state.scrollIndex + 15).map($klass => {
                    return [
                        <div key={$klass} className="classHeader">
                            <span>{$klass}</span>
                        </div>,
                        ...[1, 2, 3, 4, 5, 6].map(dow => {
                            return [1, 2, 3, 4, 5, 6, 7, 8].map(ln => {
                                let $class = 'tableCell';
                                if (ln === 1) {
                                    $class += ' ttoper';
                                }
                                let $ret = [];
                                if (this.props.timetable[$klass][dow - 1].lessons.length > (ln - 1)) {
                                    const lessn = this.props.timetable[$klass][dow - 1].lessons[ln - 1];
                                    $ret.push(<div className={$class} key={ln}>
                                        {lessn.lesson} [{lessn.kab}]
                                    </div>);
                                } else {
                                    $ret.push(<div className={$class + " tableCell_empty"} key={ln}>
                                        &mdash;
                                    </div>);
                                }

                                if (ln === 1) {
                                    $ret.unshift(<div className="ftoper"/>);
                                }
                                return $ret;
                            });
                        })
                    ];
                })}

            </div>
        );
    }
}

export default TTOutputLine;