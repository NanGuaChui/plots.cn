const jpgModules = import.meta.globEager('../assets/img/avatar/*.jpg');
const pngModules = import.meta.globEager('../assets/img/avatar/*.png');

const avatarList = {
  default: './logo.ico',
};
const modules = { ...jpgModules, ...pngModules };
for (const path in modules) {
  const userName = path.split('/').pop().split('.')[0].toLowerCase();
  avatarList[userName] = modules[path].default;
}

export default userName => {
  userName = userName.replace(/[\s,]/g, '').toLowerCase();
  if (!avatarList[userName]) {
    console.log(userName);
  }
  return avatarList[userName] || avatarList['default'] || userName;
};
