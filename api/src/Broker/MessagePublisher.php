<?php
declare(strict_types=1);

namespace App\Broker;

use App\Entity\Conversation;
use App\Entity\Message;
use Swarrot\SwarrotBundle\Broker\Publisher;
use Swarrot\Broker\Message as BrokerMessage;

class MessagePublisher
{
    /**
     * @var Publisher
     */
    private $publisher;

    /**
     * MessagePublisher constructor.
     * @param Publisher $publisher
     */
    public function __construct(Publisher $publisher)
    {
        $this->publisher = $publisher;
    }

    public function messageAdded(Message $message)
    {
        $conversation = $message->getConversation();
        foreach($conversation->getUsers() as $user) {
            if($message->getFrom() !== $user) {
                $message->setTo($user->getUsername());
                $this->publisher->publish(
                    'message',
                    new BrokerMessage(json_encode($message->denormalize())),
                    ['routing_key' => sprintf('api.conversation.%s.message.%s.added', $conversation->getId(), $message->getId())]
                );
            }
        }
    }
}
