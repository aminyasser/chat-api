class CreateApps < ActiveRecord::Migration[7.1]
  def change
    create_table :apps do |t|
      t.string :name , null: false 
      t.string :app_token , null: false , index: { unique: true }
      t.integer :chat_count , null: false 

      t.timestamps
    end

  end
end
