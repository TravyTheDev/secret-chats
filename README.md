# Secret Chats

Users can login, create a private room, send real-time invitations to other users, join rooms and chat through secure WebSockets.

Messages auto delete after 30 seconds and rooms auto delete after 10 seconds if no one is in them.

Japanese and English available.

![](https://github.com/TravyTheDev/personal-site/raw/main/public/images/private-chats.gif?raw=true)

This started mostly because everybody says to just use an auth provider when it's really not that bad to write your own and an axios interceptor. 

I also wanted to do live notifications and server side i18n on errors.

### Optimizations

The overall design is kind of bad. I don't know what's going on with the navbar. 

Right now the navbar is in multiple other components when I should have just made a layout.

I need to pay for a mail server...

### Lessons Learned

For server events users need a slice of channels because if they open up another browser it'll all hang.

Sometimes it's easier to write everything yourself.