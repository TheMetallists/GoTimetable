import React from 'react';
import '../style.scss';
import ApiHelper from '../ApiHelper';
import TTOutputLine from './TTOutputLine';

class TTableContainer extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            loading: true,
            tt: {},
            hash: ''
        };

    }

    updateDatasetIfNecessary() {
        this.setState({
            loading: true
        }, () => {
            ApiHelper.loadTimeTable((tt, hash) => {
                if (this.state.hash !== hash) {
                    this.setState({
                        loading: false,
                        tt
                    });
                }
            });
        });
    }

    componentDidMount() {
        this.updateDatasetIfNecessary();
        setInterval(() => {
            //this.updateDatasetIfNecessary();
        }, 15000);
    }

    render() {
        if (this.state.loading) {
            return <div>Loading...</div>;
        }

        return (
            <div>
                <TTOutputLine timetable={this.state.tt}/>
            </div>
        );
    }
}

export default TTableContainer;