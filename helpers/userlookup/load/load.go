package load

import (
	"github.com/audioo/goseek/helpers/ent"
)

// NoRedirSites ...
func NoRedirSites(userres string) []ent.Website {
	var arr []ent.Website

	arr = append(arr, ent.Website{Title: "FACEBOOK", Domain: "https://www.facebook.com/" + userres})
	arr = append(arr, ent.Website{Title: "BLOGSPOT", Domain: "https://" + userres + ".blogspot.com"})
	arr = append(arr, ent.Website{Title: "WORDPRESS", Domain: "https://" + userres + ".wordpress.com"})

	return arr
}

// RedirSites ...
func RedirSites(userres string) []ent.Website {
	var arr []ent.Website

	arr = append(arr, ent.Website{Title: "PINTEREST", Domain: "https://www.pinterest.com/" + userres})
	arr = append(arr, ent.Website{Title: "GITHUB", Domain: "https://www.github.com/" + userres})
	arr = append(arr, ent.Website{Title: "TUMBLR", Domain: "https://" + userres + ".tumblr.com"})
	arr = append(arr, ent.Website{Title: "FLICKR", Domain: "https://www.flickr.com/people/" + userres})
	arr = append(arr, ent.Website{Title: "VIMEO", Domain: "https://vimeo.com/" + userres})
	arr = append(arr, ent.Website{Title: "SOUNDCLOUD", Domain: "https://soundcloud.com/" + userres})
	arr = append(arr, ent.Website{Title: "DISQUS", Domain: "https://disqus.com/by/" + userres})
	arr = append(arr, ent.Website{Title: "MEDIUM", Domain: "https://medium.com/@" + userres})
	arr = append(arr, ent.Website{Title: "DEVIANTART", Domain: "https://" + userres + ".deviantart.com"})
	arr = append(arr, ent.Website{Title: "VK", Domain: "https://vk.com/" + userres})
	arr = append(arr, ent.Website{Title: "ABOUTME", Domain: "https://about.me/" + userres})
	arr = append(arr, ent.Website{Title: "FLIPBOARD", Domain: "https://flipboard.com/@" + userres})
	arr = append(arr, ent.Website{Title: "SLIDESHARE", Domain: "https://slideshare.net/" + userres})
	arr = append(arr, ent.Website{Title: "SCRIBD", Domain: "https://www.scribd.com/" + userres})
	arr = append(arr, ent.Website{Title: "PATREON", Domain: "https://www.patreon.com/" + userres})
	arr = append(arr, ent.Website{Title: "BITBUCKET", Domain: "https://bitbucket.org/" + userres})
	arr = append(arr, ent.Website{Title: "DAILYMOTION", Domain: "https://www.dailymotion.com/" + userres})
	arr = append(arr, ent.Website{Title: "ETSY", Domain: "https://www.etsy.com/shop/" + userres})
	arr = append(arr, ent.Website{Title: "CASHAPP", Domain: "https://cash.app/$" + userres})
	arr = append(arr, ent.Website{Title: "BEHANCE", Domain: "https://www.behance.net/" + userres})
	arr = append(arr, ent.Website{Title: "GOODREADS", Domain: "https://www.goodreads.com/" + userres})
	arr = append(arr, ent.Website{Title: "INSTRUCTABLES", Domain: "https://www.instructables.com/member/" + userres})
	arr = append(arr, ent.Website{Title: "KEYBASE", Domain: "https://keybase.io/" + userres})
	arr = append(arr, ent.Website{Title: "KONGREGATE", Domain: "https://kongregate.com/accounts/" + userres})
	arr = append(arr, ent.Website{Title: "LIVEJOURNAL", Domain: "https://" + userres + ".livejournal.com"})
	arr = append(arr, ent.Website{Title: "LASTFM", Domain: "https://last.fm/user/" + userres})
	arr = append(arr, ent.Website{Title: "DRIBBBLE", Domain: "https://dribbble.com/" + userres})
	arr = append(arr, ent.Website{Title: "CODECADEMY", Domain: "https://www.codecademy.com/profiles/" + userres})
	arr = append(arr, ent.Website{Title: "GRAVATAR", Domain: "https://en.gravatar.com/" + userres})
	arr = append(arr, ent.Website{Title: "PASTEBIN", Domain: "https://pastebin.com/u/" + userres})
	arr = append(arr, ent.Website{Title: "ROBLOX", Domain: "https://www.roblox.com/user.aspx?username=" + userres})
	arr = append(arr, ent.Website{Title: "GUMROAD", Domain: "https://www.gumroad.com/" + userres})
	arr = append(arr, ent.Website{Title: "NEWGROUNDS", Domain: "https://" + userres + ".newgrounds.com"})
	arr = append(arr, ent.Website{Title: "WATTPAD", Domain: "https://www.wattpad.com/user/" + userres})
	arr = append(arr, ent.Website{Title: "TRAKT", Domain: "https://www.trakt.tv/users/" + userres})
	arr = append(arr, ent.Website{Title: "BUZZFEED", Domain: "https://buzzfeed.com/" + userres})
	arr = append(arr, ent.Website{Title: "TRIPADVISOR", Domain: "https://www.tripadvisor.com/Profile/" + userres})
	arr = append(arr, ent.Website{Title: "HUBPAGES", Domain: "https://" + userres + ".hubpages.com"})
	arr = append(arr, ent.Website{Title: "BLIPFM", Domain: "https://blip.fm/" + userres})
	arr = append(arr, ent.Website{Title: "WIKIPEDIA", Domain: "https://en.wikipedia.org/wiki/User:" + userres})
	arr = append(arr, ent.Website{Title: "CODEMENTOR", Domain: "https://www.codementor.io/" + userres})
	arr = append(arr, ent.Website{Title: "REVERBNATION", Domain: "https://www.reverbnation.com/" + userres})
	arr = append(arr, ent.Website{Title: "DESIGNSPIRATION", Domain: "https://www.designspiration.com/" + userres})
	arr = append(arr, ent.Website{Title: "BANDCAMP", Domain: "https://www.bandcamp.com/" + userres})
	arr = append(arr, ent.Website{Title: "COLOURLOVERS", Domain: "https://www.colourlovers.com/lover/" + userres})
	arr = append(arr, ent.Website{Title: "IFTTT", Domain: "https://www.ifttt.com/p/" + userres})
	arr = append(arr, ent.Website{Title: "SLACK", Domain: "https://" + userres + ".slack.com"})
	arr = append(arr, ent.Website{Title: "ELLO", Domain: "https://ello.co/" + userres})
	arr = append(arr, ent.Website{Title: "7cups", Domain: "https://www.7cups.com/@" + userres + ""})
	arr = append(arr, ent.Website{Title: "Independent Academia", Domain: "https://independent.academia.edu/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Alik", Domain: "https://www.alik.cz/u/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Audio Jungle", Domain: "https://audiojungle.net/user/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Book Crossing", Domain: "https://www.bookcrossing.com/mybookshelf/" + userres + "/"})
	arr = append(arr, ent.Website{Title: "Buy me a Coffee", Domain: "https://buymeacoff.ee/" + userres + ""})
	arr = append(arr, ent.Website{Title: "CNET", Domain: "https://www.cnet.com/profiles/" + userres + "/"})
	arr = append(arr, ent.Website{Title: "CarbonMade", Domain: "https://" + userres + ".carbonmade.com"})
	arr = append(arr, ent.Website{Title: "Chess", Domain: "https://www.chess.com/member/" + userres + ""})
	arr = append(arr, ent.Website{Title: "CodeWars", Domain: "https://www.codewars.com/users/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Countable", Domain: "https://www.countable.us/" + userres + ""})
	arr = append(arr, ent.Website{Title: "DevTo", Domain: "https://dev.to/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Discogs", Domain: "https://www.discogs.com/user/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Eyeem", Domain: "https://www.eyeem.com/u/" + userres + ""})
	arr = append(arr, ent.Website{Title: "F3", Domain: "https://f3.cool/" + userres + "/"})
	arr = append(arr, ent.Website{Title: "Fortnite Tracker", Domain: "https://fortnitetracker.com/profile/all/" + userres + ""})
	arr = append(arr, ent.Website{Title: "FreeSound", Domain: "https://freesound.org/people/" + userres + "/"})
	arr = append(arr, ent.Website{Title: "GameSpot", Domain: "https://www.gamespot.com/profile/" + userres + "/"})
	arr = append(arr, ent.Website{Title: "Giphy", Domain: "https://giphy.com/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Gitlab", Domain: "https://gitlab.com/" + userres + ""})
	arr = append(arr, ent.Website{Title: "GuruShots", Domain: "https://gurushots.com/" + userres + "/photos"})
	arr = append(arr, ent.Website{Title: "Hackaday", Domain: "https://hackaday.io/" + userres + ""})
	arr = append(arr, ent.Website{Title: "HackerOne", Domain: "https://hackerone.com/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Houzz", Domain: "https://houzz.com/user/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Issuu", Domain: "https://issuu.com/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Itch", Domain: "https://" + userres + ".itch.io/"})
	arr = append(arr, ent.Website{Title: "Jimdosite", Domain: "https://" + userres + ".jimdosite.com"})
	arr = append(arr, ent.Website{Title: "LeetCode", Domain: "https://leetcode.com/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Letterboxd", Domain: "https://letterboxd.com/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Lichess", Domain: "https://lichess.org/@/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Memrise", Domain: "https://www.memrise.com/user/" + userres + "/"})
	arr = append(arr, ent.Website{Title: "Munzee", Domain: "https://www.munzee.com/m/" + userres + ""})
	arr = append(arr, ent.Website{Title: "My Anime List", Domain: "https://myanimelist.net/profile/" + userres + ""})
	arr = append(arr, ent.Website{Title: "My Mini Factory", Domain: "https://www.myminifactory.com/users/" + userres + ""})
	arr = append(arr, ent.Website{Title: "MySpace", Domain: "https://myspace.com/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Periscope", Domain: "https://www.periscope.tv/" + userres + "/"})
	arr = append(arr, ent.Website{Title: "Pink Bike", Domain: "https://www.pinkbike.com/u/" + userres + "/"})
	arr = append(arr, ent.Website{Title: "Pokemon Showdown", Domain: "https://pokemonshowdown.com/users/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Product Hunt", Domain: "https://www.producthunt.com/@" + userres + ""})
	arr = append(arr, ent.Website{Title: "PromoDJ", Domain: "http://promodj.com/" + userres + ""})
	arr = append(arr, ent.Website{Title: "PyPi", Domain: "https://pypi.org/user/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Rajce", Domain: "https://" + userres + ".rajce.idnes.cz/"})
	arr = append(arr, ent.Website{Title: "Rate Your Music", Domain: "https://rateyourmusic.com/~" + userres + ""})
	arr = append(arr, ent.Website{Title: "Red Bubble", Domain: "https://www.redbubble.com/people/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Repl.it", Domain: "https://repl.it/@" + userres + ""})
	arr = append(arr, ent.Website{Title: "Sbazar", Domain: "https://www.sbazar.cz/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Shitpost Bot", Domain: "https://www.shitpostbot.com/user/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Smule", Domain: "https://www.smule.com/" + userres + ""})
	arr = append(arr, ent.Website{Title: "SourceForge", Domain: "https://sourceforge.net/u/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Speedrun", Domain: "https://speedrun.com/user/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Roberts Space Industries", Domain: "https://robertsspaceindustries.com/citizens/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Tellonym", Domain: "https://tellonym.me/" + userres + ""})
	arr = append(arr, ent.Website{Title: "TikTok", Domain: "https://tiktok.com/@" + userres + ""})
	arr = append(arr, ent.Website{Title: "Trading View", Domain: "https://www.tradingview.com/u/" + userres + "/"})
	arr = append(arr, ent.Website{Title: "Ultimate Guitar", Domain: "https://ultimate-guitar.com/u/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Unsplash", Domain: "https://unsplash.com/@" + userres + ""})
	arr = append(arr, ent.Website{Title: "VSCO", Domain: "https://vsco.co/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Vero", Domain: "https://vero.co/" + userres + ""})
	arr = append(arr, ent.Website{Title: "WarriorForum", Domain: "https://www.warriorforum.com/members/" + userres + ".html"})
	arr = append(arr, ent.Website{Title: "We Heart It", Domain: "https://weheartit.com/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Windy", Domain: "https://community.windy.com/user/" + userres + ""})
	arr = append(arr, ent.Website{Title: "XBox", Domain: "https://xboxgamertag.com/search/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Zhihu", Domain: "https://www.zhihu.com/people/" + userres + ""})
	arr = append(arr, ent.Website{Title: "All My Links", Domain: "https://allmylinks.com/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Couch Surfing", Domain: "https://www.couchsurfing.com/people/" + userres + ""})
	arr = append(arr, ent.Website{Title: "DailyKOs", Domain: "https://www.dailykos.com/user/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Dating", Domain: "http://dating.ru/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Drive2", Domain: "https://www.drive2.ru/users/" + userres + ""})
	arr = append(arr, ent.Website{Title: "FL", Domain: "https://www.fl.ru/users/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Geocaching", Domain: "https://www.geocaching.com/p/default.aspx?u=" + userres + ""})
	arr = append(arr, ent.Website{Title: "Gfycat", Domain: "https://gfycat.com/@" + userres + ""})
	arr = append(arr, ent.Website{Title: "Hackster", Domain: "https://www.hackster.io/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Irecommend", Domain: "https://irecommend.ru/users/" + userres + ""})
	arr = append(arr, ent.Website{Title: "JBZD", Domain: "https://jbzd.com.pl/uzytkownik/" + userres + ""})
	arr = append(arr, ent.Website{Title: "JeuxVideo", Domain: "http://www.jeuxvideo.com/profil/" + userres + "?mode=infos"})
	arr = append(arr, ent.Website{Title: "Kwork", Domain: "https://kwork.ru/user/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Livelib", Domain: "https://www.livelib.ru/reader/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Moikrug", Domain: "https://moikrug.ru/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Nairaland", Domain: "https://www.nairaland.com/" + userres + ""})
	arr = append(arr, ent.Website{Title: "NN", Domain: "https://" + userres + ".www.nn.ru/"})
	arr = append(arr, ent.Website{Title: "Note", Domain: "https://note.com/" + userres + ""})
	arr = append(arr, ent.Website{Title: "NPMJS", Domain: "https://www.npmjs.com/~" + userres + ""})
	arr = append(arr, ent.Website{Title: "Osu", Domain: "https://osu.ppy.sh/users/" + userres + ""})
	arr = append(arr, ent.Website{Title: "Pikabu", Domain: "https://pikabu.ru/@" + userres + ""})
	arr = append(arr, ent.Website{Title: "Toster", Domain: "https://www.toster.ru/user/" + userres + "/answers"})

	return arr
}
