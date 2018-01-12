<?php
declare(strict_types=1);

namespace App\Command;

use App\Document\Conversation;
use App\Document\Message;
use Doctrine\ODM\MongoDB\DocumentManager;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

class CreateConversationCommand  extends Command
{
    /**
     * @var DocumentManager
     */
    private $documentManager;

    /**
     * CreateConversationCommand constructor.
     * @param DocumentManager $documentManager
     */
    public function __construct(DocumentManager $documentManager)
    {
        parent::__construct();
        $this->documentManager = $documentManager;
    }

    protected function configure()
    {
        $this->setName('app:create-conversation');
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
        $userId = '5a54017b5bb3eb011019f391';
        $conversation = new Conversation();
        $conversation->setUsers([$userId]);

        $message = new Message();
        $message->setFrom($userId);
        $message->setBody('hello world!');

        $conversation->setMessages([$message]);

        $this->documentManager->persist($conversation);
        $this->documentManager->flush();
    }
}
