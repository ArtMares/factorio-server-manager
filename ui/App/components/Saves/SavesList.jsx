import React from 'react';
import PropTypes from 'prop-types';
import Save from './Save.jsx';
import swal from 'sweetalert';

class SavesList extends React.Component {
   constructor(props) {
        super(props);
        this.updateSavesList = this.updateSavesList.bind(this);
        this.removeSave = this.removeSave.bind(this);
   }

    updateSavesList () {
        this.props.getSaves();
    }

    removeSave(saveName, e) {
        var self = this;
        swal({   
            title: "Are you sure?",  
            text: "Save: " + saveName + " will be deleted",   
            type: "warning",   
            showCancelButton: true,   
            confirmButtonColor: "#DD6B55",   
            confirmButtonText: "Yes, delete it!",   
            closeOnConfirm: false 
        }, 
        () => {
            $.ajax({
                url: "/api/saves/rm/" + saveName,
                dataType: "json",
                success: (resp) => {
                    if (resp.success === true) {
                        swal("Deleted!", resp.data, "success"); 
                        self.updateSavesList();
                    }
                }
            })
        });
    }

    render() {
        var savesList;
        if (this.props.saves.length === 0) {
            savesList = <tr></tr>
        } else {
            savesList = this.props.saves.map ( (save, i) => {
                return(
                    <Save
                        key={i}
                        saves={this.props.saves}
                        index={i}
                        save={save}
                        removeSave={this.removeSave}
                    />
                )
            });
            
        }

        return(
            <div className="box">
                <div className="box-header">
                    <h3 className="box-title">Save Files</h3>
                </div>
                
                <div className="box-body">

                    <div className="table-responsive">
                        <table className="table table-striped">
                            <thead>
                                <tr>
                                    <th>Filname</th>
                                    <th>Last Modified Time</th>
                                    <th>Filesize</th>
                                    <th>Download</th>
                                </tr>
                            </thead>
                            <tbody>
                                {savesList}
                            </tbody>
                        </table>        
                    </div>
                </div>
            </div>
        )
    }
}

SavesList.propTypes = {
    saves: PropTypes.array.isRequired,
    dlSave: PropTypes.func.isRequired,
    getSaves: PropTypes.func.isRequired
}

export default SavesList
