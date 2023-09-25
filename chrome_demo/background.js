chrome.browserAction.onClicked.addListener(function(tab) {
  chrome.tabs.create({
    url: "http://10.1.102.232:8080/fav_add?url=" + encodeURIComponent(tab.url) + "&title=" + encodeURIComponent(tab.title)
  });

});