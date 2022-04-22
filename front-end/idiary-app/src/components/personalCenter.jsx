import React, { Component } from "react";
import { Link } from "react-router-dom";
import { Upload, Button, message, Space } from "antd";
import { LoadingOutlined, PlusOutlined } from "@ant-design/icons";
import axios from "axios";
import "../css/personalCenter.css"
import auth from "../service/authService";
import * as userService from "../service/userService";

class PersonalCenter extends Component {
  state = {
    imageUrl: "",
    image_name: "",
    loading: false,
    user_id: null,
  }
  handleClick = () => {
    localStorage.removeItem("token");
  };
  customRequest = async (obj) => {
    const user = auth.getCurrentUser();
    const response = await userService.getinfobyemail(user);
    const user_id = response.data.data.user_id;
    this.setState({ user_id });

    const { file, onSuccess, onError } = obj;
    const suffix = file.name.split(".").reverse()[0];
    // console.log(file, "==file");
    // console.log(suffix, "===suffix");
    const values = {
      user_id: this.state.user_id,
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
      // console.log(res);

      this.setState({
        imageUrl: res.data.data.url,
        image_name: res.data.data.file_name,
      });
      onSuccess();


    });

    axios({
      method: "post",
      url: `/api/user/modifyProfile`,
      params: {
        user_id: this.state.user_id,
        user_profile: this.state.image_name,
      },
    }).then((res) => {
      message.success("Success change profile!");
    });
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

      <div class="container emp-profile">
        <form method="post">
          <div class="row">
            <div class="col-md-4">
              <div class="profile-img">
                {/* <img src={require("./1.jpg")} alt="" />
                <div class="file btn btn-lg btn-primary">
                  Change Photo
                  <input type="file" name="file" />
                </div> */}
                <Upload
                  name="file"
                  // action="www.baidu.com/api/upload/image"
                  // beforeUpload={this.beforeUpload}
                  className="avatar-uploader"
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
                  {!imageUrl ? (uploadButton) : null}
                </Upload>

              </div>
            </div>
            <div class="col-md-6">
              <div class="profile-head">
                <h5>Kshiti Ghelani</h5>
                <section class="section about-section gray-bg" id="about">
                  <div class="counter">
                    <div class="row">
                      <div class="col ">
                        <div class="count-data text-center">
                          <h6 class="count h2" data-to="500" data-speed="500">
                            500
                          </h6>
                          <p class="m-0px font-w-600">Diaries</p>
                        </div>
                      </div>
                      <div class="col col-lg-3">
                        <div class="count-data text-center">
                          <h6 class="count h2" data-to="150" data-speed="150">
                            150
                          </h6>
                          <p class="m-0px font-w-600">Follows</p>
                        </div>
                      </div>
                      <div class="col col-lg-3">
                        <div class="count-data text-center">
                          <h6 class="count h2" data-to="850" data-speed="850">
                            850
                          </h6>
                          <p class="m-0px font-w-600">Followers</p>
                        </div>
                      </div>
                    </div>
                  </div>
                </section>
                <ul class="nav nav-tabs" id="myTab" role="tablist">
                  <li class="nav-item">
                    <a
                      class="nav-link active clickable"
                      id="home-tab"
                      data-bs-toggle="tab"
                      data-bs-target="#home"
                      role="tab"
                      aria-controls="home"
                      aria-selected="true"
                    >
                      About
                    </a>
                  </li>
                  <li class="nav-item">
                    <a
                      class="nav-link clickable"
                      id="profile-tab"
                      data-bs-toggle="tab"
                      data-bs-target="#profile"
                      role="tab"
                      aria-controls="profile"
                      aria-selected="false"
                    >
                      Settings
                    </a>
                  </li>
                </ul>
              </div>
            </div>
            <div class="col-md-2">
              <input
                type="submit"
                class="profile-edit-btn"
                name="btnAddMore"
                value="Edit Profile"
              />
            </div>
          </div>
          <div class="row">
            <div class="col-md-4"></div>
            <div class="col-md-8">
              <div class="tab-content profile-tab" id="myTabContent">
                <div
                  class="tab-pane fade show active"
                  id="home"
                  role="tabpanel"
                  aria-labelledby="home-tab"
                >
                  <div class="row">
                    <div class="col-md-6">
                      <label>User Id</label>
                    </div>
                    <div class="col-md-6">
                      <p>Kshiti123</p>
                    </div>
                  </div>
                  <div class="row">
                    <div class="col-md-6">
                      <label>Name</label>
                    </div>
                    <div class="col-md-6">
                      <p>Kshiti Ghelani</p>
                    </div>
                  </div>
                  <div class="row">
                    <div class="col-md-6">
                      <label>Email</label>
                    </div>
                    <div class="col-md-6">
                      <p>kshitighelani@gmail.com</p>
                    </div>
                  </div>

                </div>
                <div
                  class="tab-pane fade"
                  id="profile"
                  role="tabpanel"
                  aria-labelledby="profile-tab"
                >
                  <div class="container_odd">
                    <div class="set_container">
                      <div className="mb-3">
                        <Link
                          to="/resetPassword"
                          class="links set_1_btn Vbtn-2"
                        >
                          Reset Password
                        </Link>
                      </div>

                      <hr class="my-4"></hr>
                      <div className="mb-3">
                        <a
                          href="#"
                          class="set_1_btn Vbtn-4"
                          data-bs-toggle="modal"
                          data-bs-target="#exampleModal"
                        >
                          Log Out
                        </a>
                        <div
                          class="modal fade"
                          id="exampleModal"
                          tabindex="-1"
                          aria-labelledby="exampleModalLabel"
                          aria-hidden="true"
                        >
                          <div class="modal-dialog">
                            <div class="modal-content">
                              <div class="modal-header">
                                <h5 class="modal-title" id="exampleModalLabel">
                                  Warning
                                </h5>
                                <button
                                  type="button"
                                  class="btn-close"
                                  data-bs-dismiss="modal"
                                  aria-label="Close"
                                ></button>
                              </div>
                              <div class="modal-body">
                                Are your sure you want to log out? Any unsaved
                                changes will be cleaned!
                              </div>
                              <div class="modal-footer">
                                <button
                                  type="button"
                                  class="btn btn-primary"
                                  data-bs-dismiss="modal"
                                  onClick={this.handleClick}
                                >
                                  <Link to="/" className="links text-light">
                                    Log Out
                                  </Link>
                                </button>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </form>
      </div>
    );
  }
}

export default PersonalCenter;
