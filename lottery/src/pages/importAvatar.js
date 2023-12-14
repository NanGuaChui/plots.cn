const jpgModules = import.meta.globEager("../assets/img/avatar/*.jpg");
const pngModules = import.meta.globEager("../assets/img/avatar/*.png");

const avatarList = {};
let x = 0;
const modules = { ...jpgModules, ...pngModules };
for (const path in modules) {
  const item = modules[path];
  const userName = path.split("/").pop().split(".")[0];
  avatarList[userName] = modules[path].default;

  x++;
}
console.log(x);
console.log(avatarList);

export default (userName) => {
  return avatarList[userName] || avatarList["default"] || userName;
};
