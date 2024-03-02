# Stuff to do

# Add chatrooms
-- Prompt user for which chat room they want to join, default of 0
-- map[int][]net.Conn will allow easy mapping to chat rooms
-- Server gets all messages, but will only send out the corresponding messages to the chatroom
-- Might be worth to send messages as json with the layout of {chatRoom: int,username: string, message: string}
-- Decode json and send to corresponding chatroom
# Add ability for a websocket api to hook into but keeping GoCat ability
-- Create a websocket, only if there is connections add an extra flag that will send the same json to the server.
-- Not sure if I want to implement a front end and just leave it up to myself in the future or maybe someone else
-- Metrics of username: messageCount? Probably use sqlite or something
