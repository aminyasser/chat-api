class CreateMessages < ActiveRecord::Migration[7.1]
  def change
    create_table :messages do |t|
      t.integer :message_number , null: false , index: true
      t.references :chat, foreign_key: true
      t.text :body

      t.timestamps
    end
  end
end
