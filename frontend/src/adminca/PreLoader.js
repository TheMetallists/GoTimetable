import React from "react";
import '../style.scss';

class PreLoader extends React.Component {
    render() {
        return (
            <div className={"preloader"}>
                <div className="lds-roller">
                    <div/>
                    <div/>
                    <div/>
                    <div/>
                    <div/>
                    <div/>
                    <div/>
                    <div/>
                </div>
            </div>
        );
    }
}

export default PreLoader;