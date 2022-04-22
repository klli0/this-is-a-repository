import React, { Component, PropTypes } from "react";
import ContentImg from "./content-img.jsx";
import "../css/listWb.css";
import moment from "moment";
import auth from "../service/authService";
import * as userService from "../service/userService";


export default class WeiBoList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      inputVal: "",
      transferVal: "",
      writeVal: "",
    };
  }

  /**
   * 渲染顶部View
   */
  _renderHeadView(data) {
    return (
      <div className="item">
        <div className="topRightView">
          <div className="nameandtime">
            <div style={{ marginright: "10px" }}>
              <img className="nick-img" src={data.user_profile} />
              <span style={{ marginLeft: "10px" }}>{data.nickname}</span>
            </div>
            {/* <div style={{marginLeft:"600px"}}>{data.sendTime}</div> */}

            <div style={{ marginLeft: "500px" }}>
              {moment(data.db_time * 1000).format("YYYY-MM-DD HH:mm:ss")}
            </div>
          </div>
          <div>
            <p style={{ marginTop: "10px" }}>{data.content}</p>
            <ContentImg contentImgUrls={data.image_list || []} />
          </div>
        </div>
      </div>
    );
  }

  like = (id, number) => {
    const arr = this.props.data.map((ele) => {
      if (ele.diary_id === id) {
        return { ...ele, vote_number: number + 1 };
      }
      return { ...ele };
    });
    this.props.update(arr);
  };

  comment = (id, number) => {
    const arr = this.props.data.map((ele) => {
      if (ele.diary_id === id) {
        var isshow = ele.isShowComment;
        return { ...ele, isShowComment: !isshow };
      }
      return { ...ele };
    });
    this.props.update(arr);
  };

  intChange = (e, id) => {
    const arr = this.props.data.map((ele) => {
      if (ele.diary_id === id) {
        return { ...ele, inputValue: e.target.value };
      }
      return { ...ele };
    });
    this.props.update(arr);
  };

  btnSure = (id, name, number) => {
    // 拿到后发送给服务端
    const arr = this.props.data.map((ele) => {
      if (ele.diary_id === id) {
        if (ele.comment_count === 0) {
          var commentlist = [{ nickname: name, content: ele.inputValue },];
        } else {
          var commentlist = [
            ...ele.comment_list,
            {
              nickname: name,
              content: ele.inputValue,
            },
          ];
        }
        return {
          ...ele,
          inputValue: "",
          comment_list: commentlist,
          comment_count: number + 1,
        };
      }
      return { ...ele };
    });
    this.props.update(arr);
  };

  transferdata = (id) => {
    const arr = this.props.data.map((ele) => {
      if (ele.diary_id === id) {
        var istransfer = ele.isTransfer;
        return {
          ...ele,
          isTransfer: !istransfer,
        };
      }
      return { ...ele };
    });
    this.props.update(arr);
  };

  transferChange = (e, id) => {
    const arr = this.props.data.map((ele) => {
      if (ele.diary_id === id) {
        return { ...ele, transferValue: e.target.value };
      }
      return { ...ele };
    });
    this.props.update(arr);
  };

  transfer = (id, number) => {
    // 增加转发数量
    const arr = this.props.data.map((ele) => {
      if (ele.diary_id === id) {
        return {
          ...ele,
          transferValue: "",
          share_count: number + 1,
        };
      }
      return { ...ele };
    });
    // console.log(arr);
    this.props.update(arr);

    //自己发送日记
  };

  writeChange = (e) => {
    this.setState({
      writeVal: e.target.value,
    });
  };

  write = () => {
    // 拿到后发送给服务端
  };

  render() {
    //渲染列表
    return (
      <div className="listRootViewStyle">
        {this.props.data.map((ele, index) => {
          return (
            <div key={index} style={{ marginTop: 20 }}>
              {this._renderHeadView(ele)}
              <hr className="hrStyle" />
              {/* footer */}
              <div className="commentViewStyle">
                <ul className="ulStyle">
                  <div className="shuxian"></div>
                  <li
                    className="liStyle"
                    onClick={() => this.like(ele.diary_id, ele.vote_count)}
                  >
                    点赞:{ele.vote_count}
                  </li>

                  <div className="shuxian"></div>
                  <li
                    className="liStyle"
                    onClick={() =>
                      this.comment(ele.diary_id, ele.comment_count)
                    }
                  >
                    评论:{ele.comment_count}
                  </li>

                  <div className="shuxian"></div>
                  <li
                    className="liStyle"
                    onClick={() => this.transferdata(ele.diary_id)}
                  >
                    转发:{ele.share_count}
                  </li>
                </ul>

                {ele.isTransfer && (
                  <div>
                    <input
                      style={{ width: "500px" }}
                      type="text"
                      onChange={(e) => this.transferChange(e, ele.diary_id)}
                      value={ele.transferValue}
                    />
                    <pre></pre>
                    <button
                      onClick={() =>
                        this.transfer(ele.diary_id, ele.share_count)
                      }
                    >
                      转发
                    </button>
                  </div>
                )}

                {ele.isShowComment && (
                  <div>
                    <input
                      style={{ width: "500px" }}
                      type="text"
                      onChange={(e) => this.intChange(e, ele.diary_id)}
                      value={ele.inputValue}
                    ></input>
                    <pre></pre>
                    <button
                      onClick={() =>
                        this.btnSure(
                          ele.diary_id,
                          ele.nickName,
                          ele.comment_count
                        )
                      }
                    >
                      评论
                    </button>
                  </div>
                )}
                {ele.isShowComment &&
                  !ele.comment_count === 0 &&
                  ele.comment_list.map((subEle, subIndex) => {
                    return (
                      <div key={subIndex}>
                        {subEle.nickname} : {subEle.content}
                      </div>
                    );
                  })}
              </div>
            </div>
          );
        })}
      </div>
    );
  }
}

WeiBoList.propTypes = {};
