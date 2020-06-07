import React, { Component } from 'react'

export class Info extends Component {
    constructor(props) {
        super(props);

        document.title = `Info - Calendar`;
    }
    render() {
        return (
            <div>
                <div className="row headerbar text-chonker">
                    <div className="col">
                        {"Simple Calendar App"}
                    </div>
                </div>
                <br />
                <div className="row">
                    <div className="col">
                        {"lorem ipsum"}<br />
                        {"Made by Kacper Wojakowski, github.com/kacpwoja"}
                    </div>
                </div>
            </div>
        );
    }
}