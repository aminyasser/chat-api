class MessagesController < ApplicationController
  def search
    if query_params[:query] == nil || !query_params.has_key?(:query) 
      render json: { status: 'error' , message: "you must add query" }, status: 400
      return
    end
    
    @query = query_params[:query]
    @chat = Chat.find_by!(app_token: params[:app_token], chat_number: params[:number])
    @messages = Message.search( @query , @chat.id)
    render json: {status: 'success' , messages: @messages } , status: 200 
  end



  def query_params
    params.permit(:query)
  end

end
