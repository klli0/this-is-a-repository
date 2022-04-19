import http from "./httpService";

const apiEndpoint = "http://120.78.134.104:8080/user/login";
export function register(user) {
  return http.post("http://120.78.134.104:8080/user/register", {
    email: user.email,
    verification_code: user.verification,
    password: user.password,
    username: user.username,
  });
}

export function login(user) {
  const apiUrl = `http://120.78.134.104:8080/user/login`;
  return http.post(apiUrl, {
    password: user.password,
    email: user.email,
  });
}

export function send_email(user) {
  return http.post("http://120.78.134.104:8080/email/verificationCode", {
    email: user.email,
  });
}

export function send_problem(user) {
  return http.post(apiEndpoint, {
    username: user.username,
    answer1: user.answer1,
    answer2: user.answer2,
    answer3: user.answer3,
  });
}

export function resetpassword(user) {
  return http.post("http://120.78.134.104:8080/user/modifyPassword", {
    email: user.email,
    new_password: user.password,
  });
}

export function getallinfo(user) {
  return http.post("http://120.78.134.104:8080/user/getAllUserInfo", {
    email: user,
  });
}

export function getinfobyemail(user) {
  return http.post("http://120.78.134.104:8080/user/getUserInfo", {
    email: user,
  });
}
