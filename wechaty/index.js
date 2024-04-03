const { log, ScanStatus, WechatyBuilder } = require("wechaty");
const { PuppetPadlocal } = require("wechaty-puppet-padlocal");
const http = require("http");

const puppet = new PuppetPadlocal({
  token: "token",
});
const bot = WechatyBuilder.build({
  name: "PadLocalDemo",
  puppet,
})
  .on("scan", async (qrcode, status) => {
    if (status === ScanStatus.Waiting && qrcode) {
      const qrcodeImageUrl = [
        "https://wechaty.js.org/qrcode/",
        encodeURIComponent(qrcode),
      ].join("");
      log.info(
        LOGPRE,
        `onScan:启动Wechaty成功 ${ScanStatus[status]}(${status})`
      );
    } else {
      log.info(LOGPRE, `onScan: ${ScanStatus[status]}(${status})`);
    }
  })

  .on("login", (user) => {
    log.info(LOGPRE, `${user} login`);
  })

  .on("logout", (user, reason) => {
    log.info(LOGPRE, `${user} logout, reason: ${reason}`);
  })

  .on("message", async (message) => {
    log.info(LOGPRE, `on message: ${message.toString()}`);
  })

  .on("room-invite", async (roomInvitation) => {
    log.info(LOGPRE, `on room-invite: ${roomInvitation}`);
  })

  .on("room-join", (room, inviteeList, inviter, date) => {
    log.info(
      LOGPRE,
      `on room-join, room:${room}, inviteeList:${inviteeList}, inviter:${inviter}, date:${date}`
    );
  })

  .on("room-leave", (room, leaverList, remover, date) => {
    log.info(
      LOGPRE,
      `on room-leave, room:${room}, leaverList:${leaverList}, remover:${remover}, date:${date}`
    );
  })

  .on("room-topic", (room, newTopic, oldTopic, changer, date) => {
    log.info(
      LOGPRE,
      `on room-topic, room:${room}, newTopic:${newTopic}, oldTopic:${oldTopic}, changer:${changer}, date:${date}`
    );
  })

  .on("friendship", (friendship) => {
    log.info(LOGPRE, `on friendship: ${friendship}`);
  })

  .on("error", (error) => {
    log.error(LOGPRE, `on error: ${error}`);
  });

bot.start().then(() => {
  log.info(LOGPRE, "启动wechaty成功！");
});
