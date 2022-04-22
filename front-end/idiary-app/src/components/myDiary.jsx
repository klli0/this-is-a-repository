import React, { Component, PropTypes } from "react";
import { Route, Routes } from "react-router-dom";
import axios from "axios";
import PersonalDiary from "./personalDiary";
import FriendsDiary from "./friendsDiary";
import auth from "../service/authService";
import * as userService from "../service/userService";
import Tab from "./common/tab";

import { Upload, Button, message, Space } from "antd";
import { LoadingOutlined, PlusOutlined } from "@ant-design/icons";
import "../css/myDiary.css";

export default class MyDiary extends Component {
  state = {
    Tabflag: "",
    writeVal: "",
    diary_type: "",
    list: [],
    imageUrl: "",
    loading: false,
    image_list: [],
    str_image_list: "",
    user_id: null,
  };
  componentDidMount = async () => {
    const user = auth.getCurrentUser();
    const response = await userService.getinfobyemail(user);
    const user_id = response.data.data.user_id;
    this.setState({ user_id });
    console.log("personall");
    console.log(this.state.user_id)
  };
  writeChange = (e) => {
    this.setState({
      writeVal: e.target.value,
    });
  };

  write = async () => {
    // 拿到后发送给服务端
    const user = auth.getCurrentUser();
    const response = await userService.getinfobyemail(user);
    const user_id = response.data.data.user_id;
    this.setState({ user_id });

    var t = new Date().getTime();
    axios({
      url: "/api/diary/addDiary",
      method: "post",
      params: {
        user_id: this.state.user_id,
        content: this.state.writeVal,
        authority: this.state.diary_type,
        image_list: this.state.str_image_list,
      },
    }).then((res) => {
      if (res.status === 200) {
        console.log("success");
      }
    });
    this.setState({
      writeVal: "",
      image_list: "",
      imageUrl: "",
    });
    message.success("Success add diary!");
  };
  // react一个生命周期

  handleChange = (event) => {
    const { value } = event.currentTarget;
    this.setState({ diary_type: value });
    // console.log(this.state.diary_type);
  };

  customRequest = async (obj) => {
    const { file, onSuccess, onError } = obj;
    const suffix = file.name.split(".").reverse()[0];
    // console.log(file, "==file");
    // console.log(suffix, "===suffix");
    const values = {
      user_id: 1,
      file,
    };
    const formData = new FormData();
    Object.keys(values).forEach((key) => {
      formData.append(key, values[key]);
    });
    axios({
      method: "post",
      url: `/api/upload/image`,
      data: formData,
    }).then((res) => {
      //console.log(res.data.data.file_name);

      this.setState({
        imageUrl: res.data.data.url,
        image_list: [...this.state.image_list, res.data.data.file_name],
      });
      this.setState({
        str_image_list: JSON.stringify(this.state.image_list),
      });

      //console.log(this.state.image_list);
      //console.log(this.state.str_image_list);
      // console.log("341423");
      onSuccess();
    });
  };

  success = () => {
    message.success("This is a success message");
  };

  render() {
    const { imageUrl, loading } = this.state;
    const uploadButton = (
      <div>
        {loading ? <LoadingOutlined /> : <PlusOutlined />}
        <div style={{ marginTop: 8 }}>Upload</div>
      </div>
    );
    return (
      <React.Fragment>
        <div className="wrap">
          <div className="says" onClick={this.all}>
            <h4>write a new diary</h4>
            <div style={{ display: "flex" }}>
              <textarea
                style={{ width: "500px", height: "150px" }}
                type="text"
                onChange={this.writeChange}
                value={this.state.writeVal}
              />
              <Upload
                name="file"
                // action="www.baidu.com/api/upload/image"
                // beforeUpload={this.beforeUpload}
                customRequest={this.customRequest}
                // listType="picture-card"
                showUploadList={false}
              // accept=".jpg.png.jpeg"
              // onRemove={this.onRemove}
              // showUploadList={{ showPreviewIcon: false, showDownloadIcon: false }}
              >
                {imageUrl ? (
                  <img
                    src={imageUrl}
                    alt="avatar"
                    style={{ width: 100, height: 100 }}
                  />
                ) : null}
                {uploadButton}
              </Upload>
            </div>
            <div style={{ display: "flex" }}>
              <div>
                <button
                  onClick={() => this.write()}
                  style={{ marginTop: "10px" }}
                >
                  publish
                </button>
              </div>
              <div style={{ margin: "10px" }}>
                <select
                  value={this.state.diary_type}
                  onChange={this.handleChange}
                >
                  <option value="0">Public</option>
                  <option value="1">Protected</option>
                  <option value="2">Private</option>
                </select>
              </div>
            </div>
          </div>
        </div>
        {/* <div id="upload-image"></div> */}

        <Tab />
        <Routes>
          <Route path="personalDiary" element={<PersonalDiary />} />
          <Route path="friendsDiary" element={<FriendsDiary />} />
        </Routes>
      </React.Fragment>
    );
  }
}
//ReactDOM.render(<PicturesWall />, document.getElementById("upload-image"));
