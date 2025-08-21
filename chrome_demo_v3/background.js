// chrome.browserAction.onClicked.addListener(function(tab) {
//   chrome.tabs.create({
//     url: "http://10.1.102.232:8080/fav_add?url=" + encodeURIComponent(tab.url) + "&title=" + encodeURIComponent(tab.title)
//   });

// });

chrome.action.onClicked.addListener(function (tab) {
  if (!tab.url || !tab.title) {
    console.error("Tab URL or title missing!");
    return;
  }

  const favUrl = `http://10.1.102.232:8080/fav_add?url=${encodeURIComponent(tab.url)}&title=${encodeURIComponent(tab.title)}`;
  
  chrome.tabs.create({ url: favUrl }, function () {
    console.log(`Created new tab with URL: ${favUrl}`);
  });
});