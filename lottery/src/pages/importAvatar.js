const jpgModules = import.meta.globEager('../assets/img/avatar/*.jpg');
const pngModules = import.meta.globEager('../assets/img/avatar/*.png');

const avatarList = {
  default: './logo.ico',
};
const modules = { ...jpgModules, ...pngModules };
for (const path in modules) {
  const item = modules[path];
  const userName = path.split('/').pop().split('.')[0];
  avatarList[userName] = modules[path].default;
}

export default userName => {
  return avatarList[userName] || avatarList['default'] || userName;
};
